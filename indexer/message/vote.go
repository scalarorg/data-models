package message

type VoteRequestMsg struct {
	Type    string         `gorm:"type:varchar(128)" json:"type"`
	Sender  string         `gorm:"type:varchar(64)" json:"sender"`
	PollID  string         `gorm:"type:varchar(64)" json:"poll_id"`
	Vote    map[string]any `gorm:"type:jsonb" json:"vote"`
	VoteMsg map[string]any `gorm:"type:jsonb" json:"vote_msg"`
}
