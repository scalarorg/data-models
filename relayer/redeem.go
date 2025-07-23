package relayer

import (
	"gorm.io/gorm"
)

// RedeemBlock represents the processing status of each Redeem block in the relayer
// type RedeemBlock struct {
// 	gorm.Model
// 	BlockNumber      uint64     `gorm:"type:bigint;index:idx_redeem_block_number,unique"`
// 	BlockHash        string     `gorm:"type:varchar(255)"`
// 	Chain            string     `gorm:"type:varchar(64)"`
// 	Status           string     `gorm:"type:varchar(32);default:'pending'"` // pending, processing, completed, failed
// 	ProcessedAt      *time.Time `gorm:"type:timestamp"`
// 	CompletedAt      *time.Time `gorm:"type:timestamp"`
// 	Error            string     `gorm:"type:text"`
// 	TransactionCount int        `gorm:"type:int;default:0"`
// 	ProcessedTxCount int        `gorm:"type:int;default:0"` // Number of transactions that have been processed
// }

type EvmRedeemTx struct {
	gorm.Model
	TxHash      string `gorm:"type:varchar(64);index:idx_evm_redeem_txes_tx_hash_log_index,unique"`
	LogIndex    uint64 `gorm:"type:bigint;index:idx_evm_redeem_txes_tx_hash_log_index,unique"`
	BlockNumber uint64 `gorm:"type:bigint"`
	Chain       string `gorm:"type:varchar(64)"`
	Status      string `gorm:"type:varchar(32);default:'pending'"` // pending, processing, completed, failed, failed_to_process
}
