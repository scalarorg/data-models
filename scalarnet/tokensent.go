package scalarnet

import (
	"time"

	"gorm.io/gorm"
)

type TokenSentApproved struct {
	EventID            string `gorm:"primaryKey;type:varchar(255)"`
	SourceChain        string `gorm:"type:varchar(255)"`
	SourceAddress      string `gorm:"type:varchar(255)"`
	DestinationChain   string `gorm:"type:varchar(255)"`
	DestinationAddress string `gorm:"type:varchar(255)"`
	TxHash             string `gorm:"type:varchar(255)"`
	BlockNumber        uint64
	LogIndex           uint
	Amount             uint64 `gorm:"type:bigint"`
	Symbol             string `gorm:"type:varchar(255)"`
	Status             string `gorm:"type:varchar(32);default:approved"`
	ContractAddress    string `gorm:"type:varchar(255)"`
	SourceTxHash       string `gorm:"type:varchar(255)"`
	SourceEventIndex   uint64
	CommandID          string
	TransferID         uint64    `gorm:"type:bigint"`
	CreatedAt          time.Time `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	UpdatedAt          time.Time `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	DeletedAt          gorm.DeletedAt
}
