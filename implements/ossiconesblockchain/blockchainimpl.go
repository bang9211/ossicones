package ossiconesblockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"

	"github.com/bang9211/ossicones/interfaces/blockchain"
)

var obc *ossiconesBlockchain
var once sync.Once

type block struct {
	data     string
	hash     string
	prevHash string
}

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

type ossiconesBlockchain struct {
	blocks []*block
}

func ObtainBlockchain() blockchain.Blockchain {
	if obc == nil {
		once.Do(func() {
			obc = &ossiconesBlockchain{}
			obc.AddBlock("Genesis Block")
		})
	}
	return obc
}

func (bc *ossiconesBlockchain) createBlock(data string) *block {
	newBlock := block{
		data:     data,
		prevHash: bc.getLastBlockHash(),
	}
	newBlock.calculateHash()
	return &newBlock
}

func (bc *ossiconesBlockchain) getLastBlockHash() string {
	if len(bc.blocks) > 0 {
		return bc.blocks[len(bc.blocks)-1].hash
	}
	return ""
}

func (bc *ossiconesBlockchain) AddBlock(data string) {
	newBlock := bc.createBlock(data)
	newBlock.calculateHash()
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *ossiconesBlockchain) PrintBlock() {
	for i, block := range bc.blocks {
		fmt.Println(i, ":", *block)
	}
}
