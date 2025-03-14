package event

type ChainEventConfirmed struct {
	BaseEvent
	Chain   string `gorm:"type:varchar(32)" json:"chain"`
	EventID string `gorm:"type:varchar(128)" json:"event_id"`
	Type    string `gorm:"type:varchar(128)" json:"type"`
}

type ChainEventCompleted struct {
	BaseEvent
	Chain   string `gorm:"type:varchar(32)" json:"chain"`
	EventID string `gorm:"type:varchar(128)" json:"event_id"`
	Type    string `gorm:"type:varchar(128)" json:"type"`
}

type ChainEventFailed struct {
	BaseEvent
	Chain   string `gorm:"type:varchar(32)" json:"chain"`
	EventID string `gorm:"type:varchar(128)" json:"event_id"`
	Type    string `gorm:"type:varchar(128)" json:"type"`
}

type ChainEventRetryFailed struct {
	BaseEvent
	Chain   string `gorm:"type:varchar(32)" json:"chain"`
	EventID string `gorm:"type:varchar(128)" json:"event_id"`
	Type    string `gorm:"type:varchar(128)" json:"type"`
}

type CommandBatchSigned struct {
	BaseEvent
	Chain          string `gorm:"type:varchar(32)" json:"chain"`
	CommandBatchID []byte `gorm:"type:varchar(128);serializer:HexSerializer" json:"command_batch_id"`
}

type CommandBatchAborted struct {
	BaseEvent
	Chain          string `gorm:"type:varchar(32)" json:"chain"`
	CommandBatchID []byte `gorm:"type:varchar(128);serializer:HexSerializer" json:"command_batch_id"`
}

type ConfirmTokenStarted struct {
	BaseEvent
	Chain   string `gorm:"type:varchar(32)" json:"chain"`
	EventID string `gorm:"type:varchar(128)" json:"event_id"`
	Type    string `gorm:"type:varchar(128)" json:"type"`
}

type ConfirmDepositStarted struct {
	BaseEvent
	TxID             string   `gorm:"type:varchar(128)" json:"tx_id"`
	Chain            string   `gorm:"type:varchar(32)" json:"chain"`
	DepositAddress   string   `gorm:"type:varchar(128)" json:"deposit_address"`
	TokenAddress     string   `gorm:"type:varchar(128)" json:"token_address"`
	Confirmation     uint64   `gorm:"type:bigint" json:"confirmation"`
	PollParticipants []string `gorm:"type:varchar(128)[]" json:"poll_participants"`
	Asset            Asset    `gorm:"-" json:"asset"`
	AssetDenom       string   `gorm:"type:varchar(16)" json:"asset_denom"`
	AssetAmount      uint64   `gorm:"type:bigint" json:"asset_amount"`
}

type ConfirmKeyTransferStarted struct {
	BaseEvent
	Chain              string   `gorm:"type:varchar(32)" json:"chain"`
	TxID               string   `gorm:"type:varchar(128)" json:"tx_id"`
	GatewayAddress     string   `gorm:"type:varchar(128)" json:"gateway_address"`
	ConfirmationHeight uint64   `gorm:"type:bigint" json:"confirmation_height"`
	PollParticipants   []string `gorm:"type:varchar(128)[]" json:"poll_participants"`
}

type EventConfirmSourceTxs struct {
	BaseEvent
	Chain              string   `gorm:"type:varchar(32)" json:"chain"`
	ConfirmationHeight uint64   `gorm:"type:bigint" json:"confirmation_height"`
	PoolMappings       []string `gorm:"type:varchar(128)" json:"pool_mappings"`
	Participants       []string `gorm:"type:varchar(128)[]" json:"participants"`
}

type EventTokenSent struct {
	BaseEvent
	Asset              Asset  `gorm:"-" json:"asset"`
	AssetDenom         string `gorm:"type:varchar(16)" json:"asset_denom"`
	AssetAmount        uint64 `gorm:"type:bigint" json:"asset_amount"`
	Chain              string `gorm:"type:varchar(32)" json:"chain"`
	EventID            string `gorm:"type:varchar(128)" json:"event_id"`
	TransferID         string `gorm:"type:varchar(128)" json:"transfer_id"`
	CommandID          string `gorm:"type:varchar(128)" json:"command_id"`
	Sender             string `gorm:"type:varchar(128)" json:"sender"`
	DestinationChain   string `gorm:"type:varchar(32)" json:"destination_chain"`
	DestinationAddress string `gorm:"type:varchar(128)" json:"destination_address"`
}

type EventContractCallApproved struct {
	BaseEvent
	Chain            string `gorm:"type:varchar(32)" json:"chain"`
	EventID          string `gorm:"type:varchar(128)" json:"event_id"`
	CommandID        string `gorm:"type:varchar(128)" json:"command_id"`
	Sender           string `gorm:"type:varchar(128)" json:"sender"`
	DestinationChain string `gorm:"type:varchar(32)" json:"destination_chain"`
	ContractAddress  string `gorm:"type:varchar(128)" json:"contract_address"`
	PayloadHash      string `gorm:"type:varchar(128)" json:"payload_hash"`
}

type EventContractCallWithMintApproved struct {
	BaseEvent
	Chain            string `gorm:"type:varchar(32)" json:"chain"`
	EventID          string `gorm:"type:varchar(128)" json:"event_id"`
	CommandID        []byte `gorm:"type:varchar(128);serializer:HexSerializer" json:"command_id"`
	Sender           string `gorm:"type:varchar(128)" json:"sender"`
	DestinationChain string `gorm:"type:varchar(32)" json:"destination_chain"`
	ContractAddress  string `gorm:"type:varchar(128)" json:"contract_address"`
	PayloadHash      []byte `gorm:"type:varchar(128);serializer:HexSerializer" json:"payload_hash"`
	Asset            Asset  `gorm:"-" json:"asset"`
	AssetDenom       string `gorm:"type:varchar(16)" json:"asset_denom"`
	AssetAmount      uint64 `gorm:"type:bigint" json:"asset_amount"`
}
type ContractCallFailed struct {
	BaseEvent
	Chain     string `gorm:"type:varchar(32)" json:"chain"`
	MessageID string `gorm:"type:varchar(128)" json:"message_id"`
}

type MintCommand struct {
	BaseEvent
	Chain              string `gorm:"type:varchar(32)" json:"chain"`
	TransferID         string `gorm:"type:varchar(128)" json:"transfer_id"`
	CommandID          string `gorm:"type:varchar(128)" json:"command_id"`
	DestinationChain   string `gorm:"type:varchar(32)" json:"destination_chain"`
	DestinationAddress string `gorm:"type:varchar(128)" json:"destination_address"`
	Asset              Asset  `gorm:"-" json:"asset"`
	AssetDenom         string `gorm:"type:varchar(16)" json:"asset_denom"`
	AssetAmount        uint64 `gorm:"type:bigint" json:"asset_amount"`
}

type BurnCommand struct {
	BaseEvent
	Chain              string `gorm:"type:varchar(32)" json:"chain"`
	CommandID          string `gorm:"type:varchar(128)" json:"command_id"`
	DestinationChain   string `gorm:"type:varchar(32)" json:"destination_chain"`
	DestinationAddress string `gorm:"type:varchar(128)" json:"destination_address"`
	Asset              Asset  `gorm:"-" json:"asset"`
	AssetDenom         string `gorm:"type:varchar(16)" json:"asset_denom"`
	AssetAmount        uint64 `gorm:"type:bigint" json:"asset_amount"`
}

type PollCompleted struct {
	BaseEvent
	TxID   []byte `gorm:"type:varchar(128);serializer:HexSerializer" json:"tx_id"`
	Chain  string `gorm:"type:varchar(32)" json:"chain"`
	PollID string `gorm:"type:varchar(128)" json:"poll_id"`
}
type PollExpired struct {
	BaseEvent
	TxID   []byte `gorm:"type:varchar(128);serializer:HexSerializer" json:"tx_id"`
	Chain  string `gorm:"type:varchar(32)" json:"chain"`
	PollID string `gorm:"type:varchar(128)" json:"poll_id"`
}
type PollFailed struct {
	BaseEvent
	TxID   []byte `gorm:"type:varchar(128);serializer:HexSerializer" json:"tx_id"`
	Chain  string `gorm:"type:varchar(32)" json:"chain"`
	PollID string `gorm:"type:varchar(128)" json:"poll_id"`
}

type NoEventConfirmed struct {
	BaseEvent
	TxID   []byte `gorm:"type:varchar(128);serializer:HexSerializer" json:"tx_id"`
	Chain  string `gorm:"type:varchar(32)" json:"chain"`
	PollID string `gorm:"type:varchar(128)" json:"poll_id"`
}

type TokenConfirmation struct {
	BaseEvent
	Module       string `gorm:"type:varchar(128)" json:"module"`
	Chain        string `gorm:"type:varchar(32)" json:"chain"`
	Asset        string `gorm:"-" json:"asset"`
	Symbol       string `gorm:"type:varchar(16)" json:"symbol"`
	TokenAddress string `gorm:"type:varchar(128)" json:"tokenAddress"`
	TxID         string `gorm:"type:varchar(128)" json:"txID"`
	EventID      string `gorm:"type:varchar(128)" json:"eventID"`
	Action       string `gorm:"type:varchar(128)" json:"action"`
}
