package scalarnet

import (
	"time"

	"gorm.io/gorm"
)

// Store last received events from external network
type EventCheckPoint struct {
	gorm.Model
	ChainName   string `gorm:"uniqueIndex:idx_chain_event;type:varchar(255)"`
	EventName   string `gorm:"uniqueIndex:idx_chain_event;type:varchar(255)"`
	BlockNumber uint64 `gorm:"type:bigint"`
	TxHash      string `gorm:"type:varchar(255)"`
	LogIndex    uint
	EventKey    string `gorm:"type:varchar(255)"`
}

type CallContract struct {
	EventID            string `gorm:"primaryKey;type:varchar(255)"`
	SourceChain        string `gorm:"type:varchar(255)"`
	SourceAddress      string `gorm:"type:varchar(255)"`
	DestinationChain   string `gorm:"type:varchar(255)"`
	DestinationAddress string `gorm:"type:varchar(255)"`
	TxHash             string `gorm:"type:varchar(255)"`
	BlockNumber        uint64
	LogIndex           uint
	Status             int    `gorm:"default:0"`
	SourceTxHash       string `gorm:"type:varchar(255)"`
	SourceEventIndex   uint64
	CommandId          string
	CreatedAt          time.Time `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	UpdatedAt          time.Time `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	DeletedAt          gorm.DeletedAt
}

type CallContractWithToken struct {
	CallContract
	TokenContractAddress string `gorm:"type:varchar(255)"`
	Symbol               string `gorm:"type:varchar(255)"`
	Amount               uint64 `gorm:"type:bigint"`
}

type ContractCallApprovedWithMint struct {
	ContractCallApproved
	Symbol string `gorm:"type:varchar(255)"`
	Amount uint64 `gorm:"type:bigint"`
}
type ContractCallApproved struct {
	EventID          string    `gorm:"primaryKey"`
	TxHash           string    `gorm:"type:varchar(255)"`
	SourceChain      string    `gorm:"type:varchar(255)"`
	DestinationChain string    `gorm:"type:varchar(255)"`
	CommandID        string    `gorm:"type:varchar(255)"`
	Sender           string    `gorm:"type:varchar(255)"`
	ContractAddress  string    `gorm:"type:varchar(255)"`
	PayloadHash      string    `gorm:"type:varchar(255)"`
	SourceTxHash     string    `gorm:"type:varchar(255)"`
	SourceEventIndex uint64    `gorm:"type:bigint"`
	CreatedAt        time.Time `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	UpdatedAt        time.Time `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	DeletedAt        gorm.DeletedAt
}
