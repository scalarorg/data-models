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
