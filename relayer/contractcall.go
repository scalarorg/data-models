package relayer

import (
	"gorm.io/gorm"
)

// TokenSentStatus represents the status of a token transfer
type ContractCallStatus string

const (
	// TokenSentStatusPending indicates the token is just received from the source chain
	ContractCallStatusPending ContractCallStatus = "pending"
	// TokenSentStatusVerifying indicates the transfer is broadcasting to the scalar for verification
	ContractCallStatusVerifying ContractCallStatus = "verifying"
	// TokenSentStatusApproved indicates the transfer is approved by the scalar network
	ContractCallStatusApproved ContractCallStatus = "approved"
	// TokenSentStatusSigning indicates the transfer is signing in the scalar network
	ContractCallStatusSigning ContractCallStatus = "signing"
	// TokenSentStatusExecuting indicates the transfer is executing on the destination chain
	ContractCallStatusExecuting ContractCallStatus = "executing"
	// TokenSentStatusSuccess indicates the transfer was successful executed on the destination chain
	ContractCallStatusSuccess ContractCallStatus = "success"
	// TokenSentStatusFailed indicates the transfer failed
	ContractCallStatusFailed ContractCallStatus = "failed"
	// TokenSentStatusCancelled indicates the transfer was cancelled
	ContractCallStatusCancelled ContractCallStatus = "cancelled"
)

// type ContractCallBlock struct {
// 	gorm.Model
// 	BlockNumber      uint64     `gorm:"type:bigint;index:idx_contract_call_block_number,unique"`
// 	BlockHash        string     `gorm:"type:varchar(255)"`
// 	Chain            string     `gorm:"type:varchar(64)"`
// 	Status           string     `gorm:"type:varchar(32);default:'pending'"` // pending, processing, completed, failed
// 	ProcessedAt      *time.Time `gorm:"type:timestamp"`
// 	CompletedAt      *time.Time `gorm:"type:timestamp"`
// 	Error            string     `gorm:"type:text"`
// 	TransactionCount int        `gorm:"type:int;default:0"`
// 	ProcessedTxCount int        `gorm:"type:int;default:0"` // Number of transactions that have been processed
// }

type ContractCall struct {
	gorm.Model
	BlockNumber      uint64             `gorm:"type:bigint"`
	TxHash           string             `gorm:"type:varchar(255);index:idx_contract_call_tx_hash_log_index,unique"`
	LogIndex         uint64             `gorm:"type:bigint;index:idx_contract_call_tx_hash_log_index,unique"`
	SourceChain      string             `gorm:"type:varchar(64)"`
	DestinationChain string             `gorm:"type:varchar(64)"`
	Status           ContractCallStatus `gorm:"default:pending"`
}

// type ContractCallApproved struct {
// 	EventID          string    `gorm:"primaryKey"`
// 	TxHash           string    `gorm:"type:varchar(255)"`
// 	SourceChain      string    `gorm:"type:varchar(255)"`
// 	DestinationChain string    `gorm:"type:varchar(255)"`
// 	CommandID        string    `gorm:"type:varchar(255)"`
// 	Sender           string    `gorm:"type:varchar(255)"`
// 	ContractAddress  string    `gorm:"type:varchar(255)"`
// 	PayloadHash      string    `gorm:"type:varchar(255)"`
// 	SourceTxHash     string    `gorm:"type:varchar(255)"`
// 	SourceEventIndex uint64    `gorm:"type:bigint"`
// 	CreatedAt        time.Time `gorm:"type:timestamp(6);default:current_timestamp(6)"`
// 	UpdatedAt        time.Time `gorm:"type:timestamp(6);default:current_timestamp(6)"`
// 	DeletedAt        gorm.DeletedAt
// }

type ContractCallWithToken struct {
	ContractCall
	TokenContractAddress string `gorm:"type:varchar(255)"`
	Symbol               string `gorm:"type:varchar(255)"`
	Amount               uint64 `gorm:"type:bigint"` //For redeem evm tx
}
