package message

type MultisigKeyMsg struct {
	Type   string `gorm:"type:varchar(128)" json:"type"`
	Sender string `gorm:"type:varchar(64)" json:"sender"`
	Chain  string `gorm:"type:varchar(64)" json:"chain"`
	KeyID  string `gorm:"type:varchar(64)" json:"key_id"`
}
