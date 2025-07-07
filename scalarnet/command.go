package scalarnet

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

type CommandStatus int32

const (
	CommandStatusPending     CommandStatus = 0
	CommandStatusBroadcasted CommandStatus = 1
	CommandStatusExecuted    CommandStatus = 2
	CommandStatusFailed      CommandStatus = 3
)

var CommandStatus_name = map[int32]string{
	0: "COMMAND_STATUS_PENDING",
	1: "COMMAND_STATUS_BROADCASTED",
	2: "COMMAND_STATUS_EXECUTED",
	3: "COMMAND_STATUS_FAILED",
}

var CommandStatus_value = map[string]int32{
	"COMMAND_STATUS_PENDING":     0,
	"COMMAND_STATUS_BROADCASTED": 1,
	"COMMAND_STATUS_EXECUTED":    2,
	"COMMAND_STATUS_FAILED":      3,
}

type MintCommand struct {
	gorm.Model
	TxHash           string `gorm:"type:varchar(64)"`
	SourceChain      string `gorm:"type:varchar(20);not null"`
	DestinationChain string `gorm:"type:varchar(20);not null"`
	TransferID       uint64 `gorm:"type:varchar(50);not null"`
	CommandID        string `gorm:"type:varchar(64);not null"`
	Amount           int64
	Symbol           string `gorm:"type:varchar(10);not null"`
	Recipient        string `gorm:"type:varchar(64);not null"`
}

type Command struct {
	gorm.Model
	CommandID      string `gorm:"uniqueIndex"`
	BatchCommandID string `gorm:"type:varchar(255)"`
	ChainID        string `gorm:"type:varchar(255)"`
	Params         string `gorm:"type:text"`
	KeyID          string `gorm:"type:varchar(255)"`
	CommandType    string `gorm:"type:varchar(255)"`
	Payload        []byte
	Status         CommandStatus
	ExecutedTxHash string `gorm:"type:varchar(255)"`
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
