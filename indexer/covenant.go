package indexer

type SigningPsbtStarted struct {
	BaseEvent
	Module           string `gorm:"type:varchar(128)" json:"module"`
	Chain            string `gorm:"type:varchar(32)" json:"chain"`
	SignID           uint64 `gorm:"type:bigint" json:"sign_id"`
	KeyID            string `gorm:"type:varchar(128)" json:"key_id"`
	PubKeys          string `gorm:"type:text" json:"pub_keys"`
	Psbt             string `gorm:"type:bytea" json:"psbt"`
	RequestingModule string `gorm:"type:varchar(128)" json:"requesting_module"`
}

type SigningPsbtCompleted struct {
	BaseEvent
	Module string `gorm:"type:varchar(128)" json:"module"`
	SignID uint64 `gorm:"type:bigint" json:"sign_id"`
}
type SigningPsbtExpired struct {
	BaseEvent
	Module string `gorm:"type:varchar(128)" json:"module"`
	SignID uint64 `gorm:"type:bigint" json:"sign_id"`
}
type TapScriptSigsSubmitted struct {
	BaseEvent
	Module string `gorm:"type:varchar(128)" json:"module"`
	Chain  string `gorm:"type:varchar(32)" json:"chain"`
	KeyID  string `gorm:"type:varchar(128)" json:"key_id"`
}

type KeyRotated struct {
	BaseEvent
	Module string `gorm:"type:varchar(128)" json:"module"`
	Chain  string `gorm:"type:varchar(32)" json:"chain"`
	KeyID  string `gorm:"type:varchar(128)" json:"key_id"`
}
