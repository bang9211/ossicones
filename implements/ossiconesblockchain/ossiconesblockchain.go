package ossiconesblockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"

	"github.com/bang9211/ossicones/interfaces/blockchain"
)

var obc *ossiconesBlockchain
var once sync.Once

type OssiconessBlock struct {
	Data     string
	Hash     string
	PrevHash string
}

func (b *OssiconessBlock) CalculateHash() {
	Hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", Hash)
}

type ossiconesBlockchain struct {
	blocks []*OssiconessBlock
}

func ObtainBlockchain() blockchain.Blockchain {
	if obc == nil {
		once.Do(func() {
			obc = &ossiconesBlockchain{}
			obc.AddBlock("Genesis OssiconessBlock")
		})
	}
	return obc
}

func (obc *ossiconesBlockchain) createBlock(Data string) *OssiconessBlock {
	newBlock := OssiconessBlock{
		Data:     Data,
		PrevHash: obc.getLastBlockHash(),
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

func (obc *ossiconesBlockchain) PrintBlock() {
	for i, OssiconessBlock := range obc.blocks {
		fmt.Println(i, ":", *OssiconessBlock)
	}
}
