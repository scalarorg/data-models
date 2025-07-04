package relayer

import (
	"time"

	"gorm.io/gorm"
)

// VaultBlock represents the processing status of each Vault block in the relayer
type VaultBlock struct {
	gorm.Model
	BlockNumber      uint64     `gorm:"type:bigint;index:idx_vault_block_number,unique"`
	BlockHash        string     `gorm:"type:varchar(255)"`
	Chain            string     `gorm:"type:varchar(64)"`
	Status           string     `gorm:"type:varchar(32);default:'pending'"` // pending, processing, completed, failed
	ProcessedAt      *time.Time `gorm:"type:timestamp"`
	CompletedAt      *time.Time `gorm:"type:timestamp"`
	Error            string     `gorm:"type:text"`
	TransactionCount int        `gorm:"type:int;default:0"`
	ProcessedTxCount int        `gorm:"type:int;default:0"` // Number of transactions that have been processed
}
