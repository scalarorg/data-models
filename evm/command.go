package evm

import "gorm.io/gorm"

type MintCommand struct {
	gorm.Model
	TxHash           string `gorm:"type:varchar(64)"`
	SourceChain      string `gorm:"type:varchar(20);not null"`
	DestinationChain string `gorm:"type:varchar(20);not null"`
	TransferID       uint64 `gorm:"type:varchar(50);not null"`
	CommandID        string `gorm:"type:varchar(64);not null"`
	Amount           int64
	Symbol           string `gorm:"type:varchar(10);not null"`
	Recipient        string `gorm:"type:varchar(64);not null"`
}
type CommandExecuted struct {
	gorm.Model
	ID               string `gorm:"primaryKey;type:varchar(255)"`
	SourceChain      string `gorm:"type:varchar(255)"`
	DestinationChain string `gorm:"type:varchar(255)"`
	TxHash           string `gorm:"type:varchar(255)"`
	BlockNumber      uint64
	LogIndex         uint
	CommandId        string
	Status           int     `gorm:"default:0"`
	ReferenceTxHash  *string `gorm:"type:varchar(255)"`
	Amount           *string `gorm:"type:varchar(255)"`
}
