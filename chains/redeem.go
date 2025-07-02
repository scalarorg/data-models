package chains

import (
	"time"

	"github.com/scalarorg/data-models/types"
	"gorm.io/gorm"
)

// RedeemStatus represents the status of a token transfer
type RedeemStatus string

const (
	// RedeemStatusExecuting indicates the transaction is just broadcast to the network
	RedeemStatusExecuting RedeemStatus = "executing"
	// RedeemStatusVerifying indicates the transfer is broadcasting to the scalar for verification
	RedeemStatusVerifying RedeemStatus = "verifying"
	// RedeemStatusApproved indicates the transfer is approved by the scalar network
	RedeemStatusApproved RedeemStatus = "approved"
	RedeemStatusSuccess  RedeemStatus = "success"
)

type SwitchedPhase struct {
	gorm.Model
	SourceChain       string `gorm:"type:varchar(32);index:idx_chain_txhash,unique"`
	TxHash            string `gorm:"type:varchar(255);index:idx_chain_txhash,unique"`
	BlockNumber       uint64 `gorm:"index:idx_block_number;type:bigint"`
	CustodianGroupUid string `gorm:"index:idx_custodian_group_uid;type:varchar(64)"`
	SessionSequence   uint64 `gorm:"index:session_sequence;type:bigint"`
	From              uint8  `gorm:"type:int"`
	To                uint8  `gorm:"type:int"`
}
type EvmRedeemTx struct {
	EventID              string `gorm:"primaryKey"`
	TxHash               string `gorm:"type:varchar(255)"`
	BlockNumber          uint64 `gorm:"default:0"`
	BlockTime            uint64 `gorm:"default:0"`
	LogIndex             uint
	SourceChain          string       `gorm:"type:varchar(64)"`
	SourceAddress        string       `gorm:"type:varchar(255)"`
	DestinationChain     string       `gorm:"type:varchar(64)"`
	DestinationAddress   string       `gorm:"type:varchar(255)"`
	Status               RedeemStatus `gorm:"default:pending"`
	Payload              []byte
	PayloadHash          string    `gorm:"type:varchar(255)"`
	ExecuteHash          string    `gorm:"type:varchar(255)"`
	TokenContractAddress string    `gorm:"type:varchar(255)"`
	Symbol               string    `gorm:"type:varchar(255)"`
	Amount               uint64    `gorm:"type:bigint"`
	CustodianGroupUid    string    `gorm:"type:varchar(255)"` //For redeem evm tx
	SessionSequence      uint64    `gorm:"type:bigint"`       //For redeem evm tx
	CreatedAt            time.Time `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	UpdatedAt            time.Time `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	DeletedAt            gorm.DeletedAt
}
type BtcRedeemTx struct {
	gorm.Model
	Chain             string            `gorm:"type:varchar(32)"`
	BlockNumber       uint64            `gorm:"type:bigint"`
	BlockTime         uint64            `gorm:"default:0"`
	TxHash            string            `gorm:"type:varchar(255)"`
	TxPosition        uint64            `gorm:"type:bigint"`
	RawTx             []byte            `gorm:"type:bytea"`
	MerkleProof       types.StringArray `gorm:"type:text[]"`
	CustodianGroupUid string            `gorm:"type:varchar(255)"`
	SessionSequence   uint64            `gorm:"type:bigint"`
	Symbol            string            `gorm:"type:varchar(32)"`
	TokenAddress      string            `gorm:"type:varchar(255)"`
	Amount            uint64            `gorm:"type:bigint"`
	Status            string            `gorm:"type:varchar(32)"`
}
