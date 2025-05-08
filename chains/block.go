package chains

import "gorm.io/gorm"

type BlockHeader struct {
	gorm.Model
	Chain       string `gorm:"type:varchar(255)"`
	BlockNumber uint64 `gorm:"type:bigint"`
	BlockHash   string `gorm:"type:varchar(255)"`
	BlockTime   uint64 `gorm:"type:bigint"`
}
