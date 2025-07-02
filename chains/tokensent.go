package chains

import (
	"time"

	"github.com/scalarorg/data-models/types"
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
	EventID              string            `gorm:"primaryKey;type:varchar(255)"`
	TxHash               string            `gorm:"type:varchar(255)"`
	RawTx                []byte            `gorm:"type:bytea"`
	BlockNumber          uint64            `gorm:"default:0"`
	BlockTime            uint64            `gorm:"default:0"`
	LogIndex             uint              `gorm:"default:0"`
	TxPosition           uint64            `gorm:"type:bigint"`
	MerkleProof          types.StringArray `gorm:"type:text[]"`
	SourceChain          string            `gorm:"type:varchar(64)"`
	SourceAddress        string            `gorm:"type:varchar(255)"`
	StakerPubkey         []byte            `gorm:"type:bytea"`
	DestinationChain     string            `gorm:"type:varchar(64)"`
	DestinationAddress   string            `gorm:"type:varchar(255)"`
	TokenContractAddress string            `gorm:"type:varchar(255)"`
	Amount               uint64            `gorm:"type:bigint"`
	Symbol               string            `gorm:"type:varchar(255)"`
	Status               TokenSentStatus   `gorm:"default:pending"`
	CreatedAt            time.Time         `gorm:"default:current_timestamp(6)"`
	UpdatedAt            time.Time         `gorm:"type:timestamp(6);default:current_timestamp(6)"`
}

type TokenDeployed struct {
	gorm.Model
	SourceChain  string `gorm:"type:varchar(32);uniqueIndex:idx_token_deployed_source_chain_token_address"`
	TokenAddress string `gorm:"type:varchar(255);uniqueIndex:idx_token_deployed_source_chain_token_address"`
	BlockNumber  uint64 `gorm:"type:bigint"`
	TxHash       string `gorm:"type:varchar(255)"`
	Symbol       string `gorm:"type:varchar(32)"`
}
