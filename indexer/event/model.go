package event

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

// Incase of store hex value includes prefix 0x we need 66 bytes
type BaseEvent struct {
	gorm.Model
	BlockHeight uint64    `gorm:"type:bigint" json:"block_height"`
	BlockHash   string    `gorm:"type:varchar(68)" json:"block_hash"`
	Timestamp   time.Time `json:"timestamp"`
}

type ChainAdded struct {
	BaseEvent
	Chain string `gorm:"type:varchar(32)" json:"chain"`
}

type RateLimitUpdated struct {
	BaseEvent
	Chain       string    `gorm:"type:varchar(32)" json:"chain"`
	Limit       Asset     `gorm:"-" json:"limit"`
	LimitDenom  string    `gorm:"type:varchar(16)" json:"limit_denom"`
	LimitAmount uint64    `gorm:"type:bigint" json:"limit_amount"`
	Window      time.Time `gorm:"type:interval" json:"window"`
}

type MessageExecuted struct {
	BaseEvent
	ID               string `gorm:"type:varchar(128)" json:"id"`
	SourceChain      string `gorm:"type:varchar(32)" json:"source_chain"`
	DestinationChain string `gorm:"type:varchar(32)" json:"destination_chain"`
}
