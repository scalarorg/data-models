package scalarnet

import (
	"time"

	"gorm.io/gorm"
)

type ScalarRedeemTokenApproved struct {
	EventID          string    `gorm:"primaryKey"`
	SourceChain      string    `gorm:"type:varchar(255)"`
	SourceTxHash     string    `gorm:"type:varchar(255)"`
	DestinationChain string    `gorm:"type:varchar(255)"`
	DestTxHash       string    `gorm:"type:varchar(255)"`
	CommandID        string    `gorm:"type:varchar(255)"`
	Sender           string    `gorm:"type:varchar(255)"`
	ContractAddress  string    `gorm:"type:varchar(255)"`
	PayloadHash      string    `gorm:"type:varchar(255)"`
	CreatedAt        time.Time `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	UpdatedAt        time.Time `gorm:"type:timestamp(6);default:current_timestamp(6)"`
	DeletedAt        gorm.DeletedAt
}
