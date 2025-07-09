package chains

import (
	"time"

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
	SourceChain       string `gorm:"type:varchar(32);index:idx_switched_phase,unique"`
	TxHash            string `gorm:"type:varchar(255);index:idx_switched_phase,unique"`
	BlockNumber       uint64 `gorm:"index:idx_block_number;type:bigint"`
	CustodianGroupUid string `gorm:"index:idx_switched_phase,unique;type:varchar(64)"`
	SessionSequence   uint64 `gorm:"index:idx_switched_phase,unique;type:bigint"`
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
	Chain             string `gorm:"type:varchar(64);index:idx_btc_redeem_tx_chain_tx_hash,unique"`
	BlockNumber       uint64 `gorm:"type:bigint;index:idx_btc_redeem_tx_block_number"`
	BlockHash         string `gorm:"type:varchar(255)"`
	TxHash            string `gorm:"type:varchar(255);index:idx_btc_redeem_tx_tx_hash;index:idx_btc_redeem_tx_chain_tx_hash,unique"`
	TxPosition        uint   `gorm:"type:int"`
	Amount            uint64 `gorm:"type:bigint"`
	Timestamp         uint64 `gorm:"type:bigint"`
	ServiceTag        string `gorm:"type:varchar(32)"`
	CovenantQuorum    uint8  `gorm:"type:int"`
	SessionSequence   uint64 `gorm:"type:bigint"`
	CustodianGroupUid string `gorm:"type:varchar(255)"`
	MerkleProof       []byte `gorm:"type:bytea"`
	RawTx             []byte `gorm:"type:bytea"`
}

type BtcRedeemTxOutput struct {
	gorm.Model
	Chain       string `gorm:"type:varchar(64);index:idx_btc_redeem_tx_output_chain_tx_hash,unique"`
	BlockNumber uint64 `gorm:"type:bigint;index:idx_btc_redeem_tx_output_block_number"`
	BlockHash   string `gorm:"type:varchar(255)"`
	TxHash      string `gorm:"type:varchar(255);index:idx_btc_redeem_tx_output_tx_hash;index:idx_btc_redeem_tx_output_chain_tx_hash,unique"`
	TxPosition  uint   `gorm:"type:int"`
	Amount      uint64 `gorm:"type:bigint"`
}
