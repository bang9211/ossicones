package ossiconesblockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/bang9211/ossicones/interfaces/blockchain"
)

type ossiconesBlockchain struct {
	blocks []block
}

func NewBlockchain() blockchain.Blockchain {
	return &ossiconesBlockchain{}
}

func (bc *ossiconesBlockchain) getLastBlockHash() string {
	if len(bc.blocks) > 0 {
		return bc.blocks[len(bc.blocks)-1].hash
	}
	return ""
}

func (bc *ossiconesBlockchain) AddBlock(data string) {
	newBlock := block{
		data:     data,
		hash:     "",
		prevHash: bc.getLastBlockHash(),
	}

	hash := sha256.Sum256([]byte(data + bc.getLastBlockHash()))
	newBlock.hash = fmt.Sprintf("%x", hash)

	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *ossiconesBlockchain) PrintBlock() {
	for i, block := range bc.blocks {
		fmt.Println(i, ":", block)
	}
}
