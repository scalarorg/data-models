package chains

import (
	"time"

	"gorm.io/gorm"
)

// TokenSentStatus represents the status of a token transfer
type TokenSentStatus string

const (
	// TokenSentStatusPending indicates the token is just received from the source chain
	TokenSentStatusPending TokenSentStatus = "pending"
	// TokenSentStatusVerifying indicates the transfer is broadcasting to the scalar for verification
	TokenSentStatusVerifying TokenSentStatus = "verifying"
	// TokenSentStatusApproved indicates the transfer is approved by the scalar network
	TokenSentStatusApproved TokenSentStatus = "approved"
	// TokenSentStatusSigning indicates the transfer is signing in the scalar network
	TokenSentStatusSigning TokenSentStatus = "signing"
	// TokenSentStatusExecuting indicates the transfer is executing on the destination chain
	TokenSentStatusExecuting TokenSentStatus = "executing"
	// TokenSentStatusSuccess indicates the transfer was successful executed on the destination chain
	TokenSentStatusSuccess TokenSentStatus = "success"
	// TokenSentStatusFailed indicates the transfer failed
	TokenSentStatusFailed TokenSentStatus = "failed"
	// TokenSentStatusCancelled indicates the transfer was cancelled
	TokenSentStatusCancelled TokenSentStatus = "cancelled"

	TokenSentStatusDeleted TokenSentStatus = "deleted"
)

type TokenSent struct {
	EventID   string    `gorm:"primaryKey;type:varchar(255)"`
	CreatedAt time.Time `gorm:"primaryKey;type:timestamp(6);default:current_timestamp(6)"`

	TxHash               string `gorm:"type:varchar(255)"`
	BlockNumber          uint64 `gorm:"default:0"`
	LogIndex             uint
	SourceChain          string          `gorm:"type:varchar(64)"`
	SourceAddress        string          `gorm:"type:varchar(255)"`
	DestinationChain     string          `gorm:"type:varchar(64)"`
	DestinationAddress   string          `gorm:"type:varchar(255)"`
	TokenContractAddress string          `gorm:"type:varchar(255)"`
	Amount               uint64          `gorm:"type:bigint"`
	Symbol               string          `gorm:"type:varchar(255)"`
	Status               TokenSentStatus `gorm:"default:pending"`
	UpdatedAt            time.Time       `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	DeletedAt            gorm.DeletedAt
}

type TokenDeployed struct {
	gorm.Model
	Chain        string `gorm:"type:varchar(32)"`
	BlockNumber  uint64 `gorm:"type:bigint"`
	TxHash       string `gorm:"type:varchar(255)"`
	Symbol       string `gorm:"type:varchar(32)"`
	TokenAddress string `gorm:"type:varchar(255)"`
}
