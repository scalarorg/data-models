package chains

import "gorm.io/gorm"

type CommandExecuted struct {
	gorm.Model
	SourceChain string `gorm:"type:varchar(255)"`
	Address     string `gorm:"type:varchar(255)"`
	TxHash      string `gorm:"type:varchar(255)"`
	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint
	CommandID   string `gorm:"uniqueIndex; type:varchar(255)"`
}
