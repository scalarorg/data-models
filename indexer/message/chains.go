package message

type CreateDeployTokenMsg struct {
	Type         string `gorm:"type:varchar(128)" json:"type"`
	Sender       string `gorm:"type:varchar(64)" json:"sender"`
	Chain        string `gorm:"type:varchar(64)" json:"chain"`
	TokenSymbol  string `gorm:"type:varchar(64)" json:"token_symbol"`
	AliasedToken string `gorm:"type:varchar(64)" json:"aliased_token"`
	Address      string `gorm:"type:varchar(64)" json:"address"`
}

type ConfirmTokenValueMsg struct {
	Type        string `gorm:"type:varchar(128)" json:"type"`
	Sender      string `gorm:"type:varchar(64)" json:"sender"`
	Chain       string `gorm:"type:varchar(64)" json:"chain"`
	TxID        string `gorm:"type:varchar(64)" json:"tx_id"`
	AssetChain  string `gorm:"type:varchar(64)" json:"asset_chain"`
	AssetSymbol string `gorm:"type:varchar(64)" json:"asset_symbol"`
}

type ConfirmSourceTxsMsg struct {
	Type   string   `gorm:"type:varchar(128)" json:"type"`
	Sender string   `gorm:"type:varchar(64)" json:"sender"`
	Chain  string   `gorm:"type:varchar(64)" json:"chain"`
	TxIDs  []string `gorm:"type:varchar(64)" json:"tx_ids"`
}

type SignCommandsMsg struct {
	Type   string `gorm:"type:varchar(128)" json:"type"`
	Sender string `gorm:"type:varchar(64)" json:"sender"`
	Chain  string `gorm:"type:varchar(64)" json:"chain"`
}

type CreatePendingTransfersMsg struct {
	Type   string `gorm:"type:varchar(128)" json:"type"`
	Sender string `gorm:"type:varchar(64)" json:"sender"`
	Chain  string `gorm:"type:varchar(64)" json:"chain"`
}
