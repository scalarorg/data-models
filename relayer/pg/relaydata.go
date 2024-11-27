package pg

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func (pg *PgAdapter) CreateRelayDatas(datas []RelayData, lastCheckpoint *EventCheckPoint) error {
	//Up date checkpoint and relayDatas in a transaction
	err := pg.PgClient.Transaction(func(tx *gorm.DB) error {
		result := tx.CreateInBatches(&datas, 100)
		if result.Error != nil {
			return result.Error
		}
		if lastCheckpoint != nil {
			UpdateLastEventCheckPoint(tx, lastCheckpoint)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to create relay data: %w", err)
	}
	return nil
}

func (pg *PgAdapter) updateRelayData(id string, data interface{}) (tx *gorm.DB, err error) {
	result := pg.PgClient.Model(&RelayData{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

// TODO: Find any better way to update batch relay data status
func (pg *PgAdapter) UpdateBatchRelayDataStatus(data []RelaydataExecuteResult, batchSize int) error {
	// Handle empty data case
	if len(data) == 0 {
		return nil
	}

	// Process updates in batches
	return pg.PgClient.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(data); i += batchSize {
			end := i + batchSize
			if end > len(data) {
				end = len(data)
			}

			batch := data[i:end]
			for _, item := range batch {
				updates := RelayData{
					Status: int(item.Status),
				}

				_, err := pg.updateRelayData(item.RelayDataId, updates)

				if err != nil {
					return fmt.Errorf("failed to update relay data batch: %w", err)
				}
			}
		}
		return nil
	})
}

// --- For Setup and Run Evm and Cosmos Relayer ---
func (pg *PgAdapter) UpdateRelayDataStatueWithPacketSequence(id string, status RelayDataStatus, sequence *int) error {
	data := RelayData{
		Status:         int(status),
		PacketSequence: sequence,
	}
	_, err := pg.updateRelayData(id, data)
	if err != nil {
		return err
	}
	return nil
}

func (pg *PgAdapter) UpdateRelayDataStatueWithExecuteHash(id string, status RelayDataStatus, executeHash *string) error {
	data := RelayData{
		Status:      int(status),
		ExecuteHash: executeHash,
	}
	_, err := pg.updateRelayData(id, data)
	if err != nil {
		return err
	}
	return nil
}

func (pg *PgAdapter) FindRelayDataById(id string, option *QueryOptions) (*RelayData, error) {
	var relayData RelayData

	query := pg.PgClient.Where("id = ?", id)

	// Add preload conditions based on options
	if option != nil && option.IncludeCallContract != nil && *option.IncludeCallContract {
		query = query.Preload("CallContract")
	}
	if option != nil && option.IncludeCallContractWithToken != nil && *option.IncludeCallContractWithToken {
		query = query.Preload("CallContractWithToken")
	}

	result := query.First(&relayData)
	if result.Error != nil {
		return nil, result.Error
	}

	return &relayData, nil
}

func (pg *PgAdapter) FindPayloadByHash(payloadHash string) ([]byte, error) {
	var contractCall CallContract
	result := pg.PgClient.
		Where("payload_hash = ?", strings.ToLower(payloadHash)).
		First(&contractCall)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to find payload by hash: %w", result.Error)
	}

	return contractCall.Payload, nil
}

// Find Realaydata by ContractAddress, SourceAddress, PayloadHash
// NOTE: PLEASE CHECK THE LENGTH OF RETURNED RELAYDATA IF IT IS ZERO, IT MEANS NO DATA FOUND. SINCE THIS METHOD WON'T LOG ANYTHING, PLEASE HANDLE IT IN THE CALLER.
func (pg *PgAdapter) FindRelayDataByContractCall(contractCall *CallContract) ([]RelayData, error) {
	var relayDatas []RelayData
	result := pg.PgClient.
		Joins("CallContract").
		Where("contract_address = ? AND source_address = ? AND payload_hash = ?",
			strings.ToLower(contractCall.ContractAddress),
			strings.ToLower(contractCall.SourceAddress),
			strings.ToLower(contractCall.PayloadHash)).
		Where("status IN ?", []int{int(PENDING), int(APPROVED)}).
		Preload("CallContract").
		Find(&relayDatas)

	if result.Error != nil {
		return relayDatas, fmt.Errorf("find relaydatas by contract call with error: %w", result.Error)
	}

	return relayDatas, nil
}

func preprocessOption(option *Options) error {
	if option.Size <= 0 {
		option.Size = 10
	}
	if option.Offset < 0 {
		option.Offset = 0
	}
	return nil
}

type Options struct {
	Size         int
	Offset       int
	EventId      string
	EventType    string
	EventTypes   []string
	StakerPubkey string
}

func (pg *PgAdapter) GetRelayerDatas(options *Options) ([]RelayData, int, error) {
	var relayDatas []RelayData
	var totalCount int64

	err := preprocessOption(options)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to preprocess options: %w", err)
	}

	query := pg.PgClient.Model(&RelayData{})

	// Add joins and preloads - using LEFT JOIN and ROW_NUMBER() to match original SQL
	query = query.Joins(`
		LEFT JOIN (
			SELECT 
				c.*,
				ROW_NUMBER() OVER (PARTITION BY c.id ORDER BY c.block_number) as rn
			FROM call_contracts c
		) CallContract ON relay_data.id = CallContract.id AND CallContract.rn = 1
	`).Joins(`
		LEFT JOIN (
			SELECT 
				ca.*,
				ROW_NUMBER() OVER (PARTITION BY ca.source_address, ca.contract_address, ca.payload_hash ORDER BY ca.block_number) as rn
			FROM call_contract_approveds ca
		) call_contract_approveds ON CallContract.source_address = call_contract_approveds.source_address 
			AND CallContract.contract_address = call_contract_approveds.contract_address 
			AND CallContract.payload_hash = call_contract_approveds.payload_hash 
			AND call_contract_approveds.rn = 1
	`)

	// Add conditions
	if options.EventId != "" {
		query = query.Where("relay_data.id = ?", options.EventId)
	}

	// Get total count when not searching by ID
	if options.EventId == "" {
		if err := query.Count(&totalCount).Error; err != nil {
			return nil, 0, fmt.Errorf("failed to get relay data count: %w", err)
		}
	}

	// Add pagination and ordering
	query = query.Order("relay_data.created_at DESC").
		Offset(options.Offset).
		Limit(options.Size)

	// Execute the query
	err = query.
		Preload("CallContract").
		Preload("CallContract.CallContractApproved").
		Find(&relayDatas).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get relay data: %w", err)
	}

	// For EventId searches, use the length of results as the count
	if options.EventId != "" {
		totalCount = int64(len(relayDatas))
	}

	return relayDatas, int(totalCount), nil
}

func (pg *PgAdapter) GetExecutedVaultBonding(options *Options) ([]RelayData, error) {
	var relayDatas []RelayData
	err := preprocessOption(options)
	if err != nil {
		return nil, fmt.Errorf("failed to preprocess options: %w", err)
	}

	query := pg.PgClient.Model(&RelayData{}).
		Select(`
			relay_data.*,
			call_contracts.block_number as c_block_number,
			call_contracts.tx_hash as c_tx_hash,
			call_contracts.tx_hex as c_tx_hex,
			call_contracts.log_index as c_log_index,
			call_contracts.contract_address as c_contract_address,
			call_contracts.payload as c_payload,
			call_contracts.payload_hash as c_payload_hash,
			call_contracts.source_address as c_source_address,
			call_contracts.staker_public_key as c_staker_public_key,
			call_contracts.amount as c_amount,
			command_executeds.amount as ce_amount
		`).
		Joins("JOIN call_contracts ON relay_data.id = call_contracts.id").
		Joins("LEFT JOIN command_executeds ON command_executeds.reference_tx_hash = call_contracts.tx_hash").
		Where("relay_data.status = ?", 2)

	if options.StakerPubkey != "" {
		query = query.Where("call_contracts.staker_public_key = ?", options.StakerPubkey)
	}

	query = query.Order("relay_data.created_at DESC").
		Offset(options.Offset).
		Limit(options.Size)

	var results []struct {
		RelayData
		CBlockNumber     *uint64 `gorm:"column:c_block_number"`
		CTxHash          *string `gorm:"column:c_tx_hash"`
		CTxHex           []byte  `gorm:"column:c_tx_hex"`
		CLogIndex        *uint   `gorm:"column:c_log_index"`
		CContractAddress *string `gorm:"column:c_contract_address"`
		CPayload         []byte  `gorm:"column:c_payload"`
		CPayloadHash     *string `gorm:"column:c_payload_hash"`
		CSourceAddress   *string `gorm:"column:c_source_address"`
		CStakerPublicKey *string `gorm:"column:c_staker_public_key"`
		CAmount          *uint64 `gorm:"column:c_amount"`
		CEAmount         *string `gorm:"column:ce_amount"`
	}

	if err := query.Find(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to get executed vault bonding: %w", err)
	}

	// Convert results to RelayData slice
	for _, result := range results {
		relayData := result.RelayData
		relayData.CallContract = &CallContract{
			ID:              relayData.ID,
			BlockNumber:     *result.CBlockNumber,
			TxHash:          *result.CTxHash,
			TxHex:           result.CTxHex,
			LogIndex:        *result.CLogIndex,
			ContractAddress: *result.CContractAddress,
			Payload:         result.CPayload,
			PayloadHash:     *result.CPayloadHash,
			SourceAddress:   *result.CSourceAddress,
			StakerPublicKey: result.CStakerPublicKey,
			Amount:          *result.CAmount,
			CommandExecuted: &CommandExecuted{
				Amount: result.CEAmount,
			},
		}
		relayDatas = append(relayDatas, relayData)
	}

	return relayDatas, nil
}
