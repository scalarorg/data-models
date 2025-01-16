package chains

import (
	"time"

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

type ContractCall struct {
	EventId             string `gorm:"primaryKey"`
	TxHash              string `gorm:"type:varchar(255)"`
	TxHex               []byte
	BlockNumber         uint64 `gorm:"default:0"`
	LogIndex            uint
	SourceChain         string `gorm:"type:varchar(64)"`
	SourceAddress       string `gorm:"type:varchar(255)"`
	Payload             []byte
	PayloadHash         string             `gorm:"type:varchar(64);"`
	DestinationChain    string             `gorm:"type:varchar(64)"`
	DestContractAddress string             `gorm:"type:varchar(255)"`
	StakerPublicKey     *string            `gorm:"type:varchar(255)"`
	Status              ContractCallStatus `gorm:"type:varchar(32);default:pending"`
	ExecuteHash         *string            `gorm:"type:varchar(255)"`
	CreatedAt           time.Time          `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	UpdatedAt           time.Time          `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	DeletedAt           gorm.DeletedAt
}

type ContractCallWithToken struct {
	ContractCall
	TokenContractAddress string `gorm:"type:varchar(255)"`
	Symbol               string `gorm:"type:varchar(255)"`
	Amount               uint64 `gorm:"type:bigint"`
}
