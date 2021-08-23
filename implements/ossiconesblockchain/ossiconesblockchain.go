package ossiconesblockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"

	"github.com/bang9211/ossicones/interfaces/blockchain"
)

var obc *ossiconesBlockchain
var once sync.Once

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

type ossiconesBlockchain struct {
	blocks []*OssiconesBlock
}

// GetOrCreate returns the existing singletone object of OssiconesBlockchain if present.
// Otherwise. it creates and returns the object.
func GetOrCreate() blockchain.Blockchain {
	if obc == nil {
		once.Do(func() {
			obc = &ossiconesBlockchain{}
			obc.AddBlock("Genesis OssiconesBlock")
		})
	}
	return obc
}

func (obc *ossiconesBlockchain) createBlock(Data string) *OssiconesBlock {
	newBlock := OssiconesBlock{
		Data:     Data,
		PrevHash: obc.getLastBlockHash(),
		Height:   len(obc.blocks) + 1,
	}
	newBlock.CalculateHash()
	return &newBlock
}

func (obc *ossiconesBlockchain) getLastBlockHash() string {
	if len(obc.blocks) > 0 {
		return obc.blocks[len(obc.blocks)-1].Hash
	}
	return ""
}

func (obc *ossiconesBlockchain) AddBlock(Data string) {
	newBlock := obc.createBlock(Data)
	newBlock.CalculateHash()
	obc.blocks = append(obc.blocks, newBlock)
}

func (obc *ossiconesBlockchain) AllBlocks() []interface{} {
	blocks := make([]interface{}, len(obc.blocks))
	for i, block := range obc.blocks {
		blocks[i] = block
	}
	return blocks
}

func (obc *ossiconesBlockchain) GetBlock(height int) (blockchain.Block, error) {
	if height > len(obc.blocks) {
		return nil, blockchain.ErrorNotFound
	}
	return obc.blocks[height-1], nil
}

func (obc *ossiconesBlockchain) PrintBlock() {
	for i, OssiconesBlock := range obc.blocks {
		fmt.Println(i, ":", *OssiconesBlock)
	}
}
