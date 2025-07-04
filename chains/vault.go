package chains

import (
	"encoding/binary"
	"encoding/hex"

	"gorm.io/gorm"
)

type HeaderEntry struct {
	Hex    string `json:"hex"` //Header hex
	Hash   string `json:"hash"`
	Height int    `json:"height"` //Height of the block
}
type BtcBlockHeader struct {
	Version       int32  `json:"version"` //TODO: remove this field
	PrevBlockhash []byte `json:"prev_blockhash"`
	MerkleRoot    []byte `json:"merkle_root"`
	Time          uint32 `json:"time"`
	CompactTarget uint32 `json:"compact_target"`
	Nonce         uint32 `json:"nonce"`
	Hash          string `json:"hash"`
	Height        int    `json:"height" gorm:"type:bigint;index:idx_btc_block_header_height,unique"` //Height of the block
}

func (b *BtcBlockHeader) ParseHeaderEntry(headerEntry *HeaderEntry) error {
	bytes, err := hex.DecodeString(headerEntry.Hex)
	if err != nil {
		return err
	}
	version := binary.LittleEndian.Uint32(bytes[0:4])
	time := binary.LittleEndian.Uint32(bytes[68:72])
	compactTarget := binary.LittleEndian.Uint32(bytes[72:76])
	nonce := binary.LittleEndian.Uint32(bytes[76:80])

	b.Version = int32(version)
	b.PrevBlockhash = bytes[4:36]
	b.MerkleRoot = bytes[36:68]
	b.Time = time
	b.CompactTarget = compactTarget
	b.Nonce = nonce
	b.Hash = headerEntry.Hash
	b.Height = headerEntry.Height
	return nil
}

type VaultTransaction struct {
	gorm.Model
	Chain                       string `gorm:"type:varchar(64);index:idx_vault_tx_chain_tx_hash,unique"`
	BlockNumber                 uint64 `gorm:"type:bigint;index:idx_vault_tx_block_number"`
	BlockHash                   string `gorm:"type:varchar(255)"`
	TxHash                      string `gorm:"type:varchar(255);index:idx_vault_tx_chain_tx_hash,unique"`
	TxPosition                  uint   `gorm:"type:int"`
	Amount                      uint64 `gorm:"type:bigint"`
	StakerScriptPubkey          string `gorm:"type:varchar(255)"`
	Timestamp                   uint64 `gorm:"type:bigint"`
	ChangeAmount                uint64 `gorm:"type:bigint"`
	ChangeAddress               string `gorm:"type:varchar(255)"`
	ServiceTag                  string `gorm:"type:varchar(32)"`
	CovenantQuorum              uint8  `gorm:"type:int"`
	VaultTxType                 uint8  `gorm:"type:int"` // 1=Staking, 2=Unstaking
	DestinationChain            uint64 `gorm:"type:bigint"`
	DestinationTokenAddress     string `gorm:"type:varchar(255)"`
	DestinationRecipientAddress string `gorm:"type:varchar(255)"`
	SessionSequence             uint64 `gorm:"type:bigint"`
	CustodianGroupUID           string `gorm:"type:varchar(255)"`
	ScriptPubkey                []byte `gorm:"type:bytea"`
	MerkleProof                 []byte `gorm:"type:bytea"`
	RawTx                       []byte `gorm:"type:bytea"`
}
