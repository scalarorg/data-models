package chains

import "gorm.io/gorm"

// RedeemStatus represents the status of a token transfer
type RedeemStatus string

const (
	// RedeemStatusExecuting indicates the transaction is just broadcast to the network
	RedeemStatusExecuting RedeemStatus = "excuting"
	// RedeemStatusVerifying indicates the transfer is broadcasting to the scalar for verification
	RedeemStatusVerifying RedeemStatus = "verifying"
	// RedeemStatusApproved indicates the transfer is approved by the scalar network
	RedeemStatusApproved RedeemStatus = "approved"
)

type SwitchedPhase struct {
	gorm.Model
	Chain             string `gorm:"type:varchar(32)"`
	BlockNumber       uint64 `gorm:"type:bigint"`
	TxHash            string `gorm:"type:varchar(255)"`
	CustodianGroupUid string `gorm:"type:varchar(64)"`
	SessionSequence   uint64 `gorm:"type:bigint"`
	From              uint8  `gorm:"type:int"`
	To                uint8  `gorm:"type:int"`
}

type RedeemTx struct {
	gorm.Model
	Chain             string `gorm:"type:varchar(32)"`
	BlockNumber       uint64 `gorm:"type:bigint"`
	TxHash            string `gorm:"type:varchar(255)"`
	CustodianGroupUid string `gorm:"type:varchar(255)"`
	Symbol            string `gorm:"type:varchar(32)"`
	TokenAddress      string `gorm:"type:varchar(255)"`
	Amount            uint64 `gorm:"type:bigint"`
	Status            string `gorm:"type:varchar(32)"`
}
