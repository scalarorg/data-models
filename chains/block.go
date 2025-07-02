package chains

import "gorm.io/gorm"

type BlockHeader struct {
	gorm.Model
	Chain       string `gorm:"type:varchar(64);index:idx_chain_blocknumber,unique"`
	BlockNumber uint64 `gorm:"type:bigint;index:idx_chain_blocknumber,unique"`
	ParentHash  string `gorm:"type:varchar(64)"`
	BlockHash   string `gorm:"type:varchar(64)"`
	BlockTime   uint64 `gorm:"type:bigint"`
	TxHash      string `gorm:"type:varchar(64)"`
	Root        string `gorm:"type:varchar(64)"`
	BeaconRoot  string `gorm:"type:varchar(64)"`
}
