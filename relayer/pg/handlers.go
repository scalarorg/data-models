package pg

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (pg *PgAdapter) CreateSingleValue(value any) error {
	result := pg.PgClient.Create(value)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pg *PgAdapter) CreateBatchValue(values any, batchSize int) error {
	result := pg.PgClient.CreateInBatches(values, batchSize)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pg *PgAdapter) GetLastEventCheckPoint(chainName, eventName string) (*EventCheckPoint, error) {
	//Default value
	lastBlock := EventCheckPoint{
		ChainName:   chainName,
		EventName:   eventName,
		BlockNumber: 0,
		TxHash:      "",
		LogIndex:    0,
		EventKey:    "",
	}
	result := pg.PgClient.Where("chain_name = ? AND event_name = ?", chainName, eventName).First(&lastBlock)
	return &lastBlock, result.Error
}

func (pg *PgAdapter) UpdateLastEventCheckPoint(value *EventCheckPoint) error {
	return UpdateLastEventCheckPoint(pg.PgClient, value)
}

// For transactional update
func UpdateLastEventCheckPoint(db *gorm.DB, value *EventCheckPoint) error {
	result := db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "chain_name"}, {Name: "event_name"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"block_number": value.BlockNumber,
				"tx_hash":      value.TxHash,
				"log_index":    value.LogIndex,
				"event_key":    value.EventKey,
			}),
		},
	).Create(value)
	if result.Error != nil {
		return fmt.Errorf("failed to update last event check point: %w", result.Error)
	}

	return nil
}

func (pg *PgAdapter) UpdateEventStatus(id string, status RelayDataStatus) error {
	data := RelayData{
		Status: int(status),
	}

	result := pg.PgClient.Model(&RelayData{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
