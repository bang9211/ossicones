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
	data     string
	hash     string
	prevHash string
}

func (b *OssiconessBlock) CalculateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
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

func (obc *ossiconesBlockchain) createBlock(data string) *OssiconessBlock {
	newBlock := OssiconessBlock{
		data:     data,
		prevHash: obc.getLastBlockHash(),
	}
	newBlock.CalculateHash()
	return &newBlock
}

func (obc *ossiconesBlockchain) getLastBlockHash() string {
	if len(obc.blocks) > 0 {
		return obc.blocks[len(obc.blocks)-1].hash
	}
	return ""
}

func (obc *ossiconesBlockchain) AddBlock(data string) {
	newBlock := obc.createBlock(data)
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
