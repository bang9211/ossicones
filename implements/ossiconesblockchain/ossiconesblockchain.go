package ossiconesblockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
)

const defaultGenesisBlockData = "Genesis OssiconesBlock"

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

func (b *OssiconesBlock) GetData() string {
	return b.Data
}

type ossiconesBlockchain struct {
	blocks []*OssiconesBlock
}

// GetOrCreate returns the existing singletone object of OssiconesBlockchain if present.
// Otherwise, it creates and returns the object.
func GetOrCreate(config config.Config) blockchain.Blockchain {
	if obc == nil {
		once.Do(func() {
			obc = &ossiconesBlockchain{}
			data := config.GetString(
				"OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA",
				defaultGenesisBlockData)
			obc.AddBlock(data)
		})
	}
	return obc
}

func (o *ossiconesBlockchain) createBlock(Data string) *OssiconesBlock {
	newBlock := OssiconesBlock{
		Data:     Data,
		PrevHash: o.getLastBlockHash(),
		Height:   len(o.blocks) + 1,
	}
	newBlock.CalculateHash()
	return &newBlock
}

func (o *ossiconesBlockchain) getLastBlockHash() string {
	if len(o.blocks) > 0 {
		return o.blocks[len(o.blocks)-1].Hash
	}
	return ""
}

func (o *ossiconesBlockchain) AddBlock(Data string) {
	newBlock := o.createBlock(Data)
	newBlock.CalculateHash()
	o.blocks = append(o.blocks, newBlock)
}

func (o *ossiconesBlockchain) AllBlocks() []interface{} {
	blocks := make([]interface{}, len(o.blocks))
	for i, block := range o.blocks {
		blocks[i] = block
	}
	return blocks
}

func (o *ossiconesBlockchain) GetBlock(height int) (blockchain.Block, error) {
	if height > len(o.blocks) {
		return nil, blockchain.ErrorNotFound
	}
	return o.blocks[height-1], nil
}

func (o *ossiconesBlockchain) PrintBlock() {
	for i, OssiconesBlock := range o.blocks {
		fmt.Println(i, ":", *OssiconesBlock)
	}
}

func (o *ossiconesBlockchain) Close() error {
	obc = nil
	return nil
}
