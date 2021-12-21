package ossiconesblockchain

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/bang9211/ossicones/utils"
)

// OssiconesBlock for OssiconesBlockChain.
type OssiconesBlock struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevhash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

/*
func (b *OssiconesBlock) CalculateHash() {
	Hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", Hash)
}
*/

func (b *OssiconesBlock) GetData() string {
	return b.Data
}

func (b *OssiconesBlock) persist() error {
	block, err := utils.ToBytes(b)
	if err != nil {
		return err
	}
	return db.SaveBlock(b.Hash, block)
}

func (o *OssiconesBlock) Mine() {
	target := strings.Repeat("0", defaultDifficulty)
	for {
		byteBlock := []byte(fmt.Sprint(o))
		hash := fmt.Sprintf("%x", sha256.Sum256(byteBlock))
		if strings.HasPrefix(hash, target) {
			o.Hash = hash
			return
		} else {
			o.Nonce++
		}
	}
}

func (b *OssiconesBlock) restore(data []byte) error {
	return utils.FromBytes(b, data)
}

func (b *OssiconesBlock) GetPrevHash() string {
	return b.PrevHash
}
