package message

type BatchRequestMsg struct {
	Type     string   `gorm:"type:varchar(128)" json:"type"`
	Sender   string   `gorm:"type:varchar(64)" json:"sender"`
	Messages []string `gorm:"type:json" json:"messages"`
}
