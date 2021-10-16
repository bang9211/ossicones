package ossiconesblockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/bang9211/ossicones/utils"
)

// OssiconesBlock for OssiconesBlockChain.
type OssiconesBlock struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevhash,omitempty"`
	Height   int    `json:"height"`
}

func (b *OssiconesBlock) CalculateHash() {
	Hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", Hash)
}

func (b *OssiconesBlock) GetData() string {
	return b.Data
}

func (b *OssiconesBlock) persist() error {
	block, err := utils.ToBytes(b)
	if err != nil {
		return err
	}
	db.SaveBlock(b.Hash, block)
	return nil
}

func (b *OssiconesBlock) restore(data []byte) error {
	return utils.FromBytes(b, data)
}

func (b *OssiconesBlock) GetPrevHash() string {
	return b.PrevHash
}
