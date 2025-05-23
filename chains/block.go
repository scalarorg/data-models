package chains

import "gorm.io/gorm"

type BlockHeader struct {
	gorm.Model
	Chain       string `gorm:"type:varchar(255);index:idx_chain_blocknum,unique"`
	BlockNumber uint64 `gorm:"type:bigint;index:idx_chain_blocknum,unique"`
	BlockHash   string `gorm:"type:varchar(255)"`
	BlockTime   uint64 `gorm:"type:bigint"`
}
