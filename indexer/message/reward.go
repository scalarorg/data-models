package message

type RefundMsg struct {
	Type         string         `gorm:"type:varchar(128)" json:"@type,omitempty"`
	Sender       string         `gorm:"type:varchar(64)" json:"sender,omitempty"`
	InnerMessage map[string]any `gorm:"type:jsonb" json:"inner_message,omitempty"`
}
