package relayer

type BlockStatus string

const (
	BlockStatusPending    BlockStatus = "pending"
	BlockStatusProcessing BlockStatus = "processing"
	BlockStatusCompleted  BlockStatus = "completed"
	BlockStatusFailed     BlockStatus = "failed"
)
