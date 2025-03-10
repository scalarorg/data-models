package event

type Voted struct {
	BaseEvent
	Module string `gorm:"type:varchar(128)" json:"module"`
	Action string `gorm:"type:varchar(64)" json:"action"`
	PollID string `gorm:"type:varchar(128)" json:"poll"`
	Voter  string `gorm:"type:varchar(128)" json:"voter"`
	State  string `gorm:"type:varchar(128)" json:"state"`
}
