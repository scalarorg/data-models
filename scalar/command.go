package scalar

import (
	"gorm.io/gorm"
)

type BatchedCommandsStatus int32

const (
	BatchNonExistent BatchedCommandsStatus = 0
	BatchSigning     BatchedCommandsStatus = 1
	BatchAborted     BatchedCommandsStatus = 2
	BatchSigned      BatchedCommandsStatus = 3
)

var BatchedCommandsStatus_name = map[int32]string{
	0: "BATCHED_COMMANDS_STATUS_UNSPECIFIED",
	1: "BATCHED_COMMANDS_STATUS_SIGNING",
	2: "BATCHED_COMMANDS_STATUS_ABORTED",
	3: "BATCHED_COMMANDS_STATUS_SIGNED",
}

var BatchedCommandsStatus_value = map[string]int32{
	"BATCHED_COMMANDS_STATUS_UNSPECIFIED": 0,
	"BATCHED_COMMANDS_STATUS_SIGNING":     1,
	"BATCHED_COMMANDS_STATUS_ABORTED":     2,
	"BATCHED_COMMANDS_STATUS_SIGNED":      3,
}

type Command struct {
	gorm.Model
	CommandID      string `gorm:"type:varchar(255)"`
	BatchCommandID []byte
	Params         []byte
	KeyID          string `gorm:"type:varchar(255)"`
	CommandType    int
	Payload        []byte
}
type BatchCommand struct {
	gorm.Model
	BatchCommandID        []byte
	Data                  []byte
	SigHash               []byte
	Status                BatchedCommandsStatus
	KeyID                 string `gorm:"type:varchar(255)"`
	PrevBatchedCommandsID []byte
}
