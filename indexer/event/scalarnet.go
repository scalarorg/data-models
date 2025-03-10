package event

type ContractCallSubmitted struct {
	BaseEvent
	MessageID        string `gorm:"type:varchar(128)" json:"message_id"`
	Sender           string `gorm:"type:varchar(128)" json:"sender"`
	SourceChain      string `gorm:"type:varchar(32)" json:"source_chain"`
	DestinationChain string `gorm:"type:varchar(32)" json:"destination_chain"`
	ContractAddress  string `gorm:"type:varchar(128)" json:"contract_address"`
	PayloadHash      string `gorm:"type:varchar(128)" json:"payload_hash"`
	Payload          string `gorm:"type:bytea" json:"payload"`
}
type ContractCallWithTokenSubmitted struct {
	BaseEvent
	MessageID        string `gorm:"type:varchar(128)" json:"message_id"`
	Sender           string `gorm:"type:varchar(128)" json:"sender"`
	SourceChain      string `gorm:"type:varchar(32)" json:"source_chain"`
	DestinationChain string `gorm:"type:varchar(32)" json:"destination_chain"`
	ContractAddress  string `gorm:"type:varchar(128)" json:"contract_address"`
	PayloadHash      string `gorm:"type:varchar(128)" json:"payload_hash"`
	Payload          string `gorm:"type:bytea" json:"payload"`
	Asset            Asset  `gorm:"-" json:"asset"`
	AssetDenom       string `gorm:"type:varchar(16)" json:"asset_denom"`
	AssetAmount      uint64 `gorm:"type:bigint" json:"asset_amount"`
}
type TokenSent struct {
	BaseEvent
	TransferID       string `gorm:"type:varchar(128)" json:"transfer_id"`
	Sender           string `gorm:"type:varchar(128)" json:"sender"`
	SourceChain      string `gorm:"type:varchar(32)" json:"source_chain"`
	DestinationChain string `gorm:"type:varchar(32)" json:"destination_chain"`
	Asset            Asset  `gorm:"-" json:"asset"`
	AssetDenom       string `gorm:"type:varchar(16)" json:"asset_denom"`
	AssetAmount      uint64 `gorm:"type:bigint" json:"asset_amount"`
}

type FeePaid struct {
	BaseEvent
	MessageID        string `gorm:"type:varchar(128)" json:"message_id"`
	Recipient        string `gorm:"type:varchar(128)" json:"recipient"`
	Fee              Asset  `gorm:"-" json:"fee"`
	FeeDenom         string `gorm:"type:varchar(16)" json:"fee_denom"`
	FeeAmount        uint64 `gorm:"type:bigint" json:"fee_amount"`
	Asset            string `gorm:"type:varchar(16)" json:"asset"`
	SourceChain      string `gorm:"type:varchar(32)" json:"source_chain"`
	DestinationChain string `gorm:"type:varchar(32)" json:"destination_chain"`
}
type ScalarTransferCompleted struct {
	BaseEvent
	ID          string `gorm:"type:varchar(128)" json:"transfer_id"`
	Recipient   string `gorm:"type:varchar(128)" json:"recipient"`
	Asset       Asset  `gorm:"-" json:"asset"`
	AssetDenom  string `gorm:"type:varchar(16)" json:"asset_denom"`
	AssetAmount uint64 `gorm:"type:bigint" json:"asset_amount"`
}
type FeeCollected struct {
	BaseEvent
	Collector string `gorm:"type:varchar(128)" json:"collector"`
	Fee       Asset  `gorm:"-" json:"fee"`
	FeeDenom  string `gorm:"type:varchar(16)" json:"fee_denom"`
	FeeAmount uint64 `gorm:"type:bigint" json:"fee_amount"`
}
