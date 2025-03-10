package event

type KeygenStarted struct {
	BaseEvent
	Module       string   `gorm:"type:varchar(128)" json:"module"`
	KeyID        string   `gorm:"type:varchar(128)" json:"key_id"`
	Participants []string `gorm:"type:varchar(128)[]" json:"participants"`
}
type KeygenCompleted struct {
	BaseEvent
	Module string `gorm:"type:varchar(128)" json:"module"`
	KeyID  string `gorm:"type:varchar(128)" json:"key_id"`
}
type KeygenExpired struct {
	BaseEvent
	Module string `gorm:"type:varchar(128)" json:"module"`
	KeyID  string `gorm:"type:varchar(128)" json:"key_id"`
}
type KeyAssigned struct {
	BaseEvent
	Module string `gorm:"type:varchar(128)" json:"module"`
	Chain  string `gorm:"type:varchar(32)" json:"chain"`
	KeyID  string `gorm:"type:varchar(128)" json:"key_id"`
}
type MultiSigKeyRotated struct {
	BaseEvent
	Module string `gorm:"type:varchar(128)" json:"module"`
	Chain  string `gorm:"type:varchar(32)" json:"chain"`
	KeyID  string `gorm:"type:varchar(128)" json:"key_id"`
}
type SigningStarted struct {
	BaseEvent
	Module           string `gorm:"type:varchar(128)" json:"module"`
	SignID           uint64 `gorm:"type:bigint" json:"sign_id"`
	KeyID            string `gorm:"type:varchar(128)" json:"key_id"`
	PubKeys          string `gorm:"type:text" json:"pub_keys"`
	PayloadHash      string `gorm:"type:varchar(128)" json:"payload_hash"`
	RequestingModule string `gorm:"type:varchar(128)" json:"requesting_module"`
}

type SigningCompleted struct {
	BaseEvent
	Module string `gorm:"type:varchar(128)" json:"module"`
	SignID uint64 `gorm:"type:bigint" json:"sign_id"`
}

type SigningExpired struct {
	BaseEvent
	Module string `gorm:"type:varchar(128)" json:"module"`
	SignID uint64 `gorm:"type:bigint" json:"sign_id"`
}

type SignatureSubmitted struct {
	BaseEvent
	Module       string   `gorm:"type:varchar(128)" json:"module"`
	SignID       uint64   `gorm:"type:bigint" json:"sign_id"`
	Participants []string `gorm:"type:text[]" json:"participants"`
	Signature    string   `gorm:"type:bytea" json:"signature"`
}

type PubkeySubmitted struct {
	BaseEvent
	Module       string   `gorm:"type:varchar(128)" json:"module"`
	KeyID        string   `gorm:"type:varchar(128)" json:"key_id"`
	Participants []string `gorm:"type:text[]" json:"participants"`
	Pubkey       string   `gorm:"type:text" json:"pubkey"`
}

type KeygenOptIn struct {
	BaseEvent
	Participant string `gorm:"type:varchar(128)" json:"participant"`
}

type KeygenOptOut struct {
	BaseEvent
	Participant string `gorm:"type:varchar(128)" json:"participant"`
}
