package event

type Heartbeat struct {
	BaseEvent
	Module string `gorm:"type:varchar(128)" json:"module"`
	Action string `gorm:"type:varchar(128)" json:"action"`
}

type Message struct {
	BaseEvent
}
