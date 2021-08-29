package blockchain

import "errors"

type Block interface {
	// CalculateHash calculates hash using sha256.
	CalculateHash()
	// GetData gets data of the block.
	GetData() string
}

type Blockchain interface {
	// AddBlock adds data to blockchain.
	AddBlock(data string)
	// AllBlocks gets all the blocks of this blockchain.
	AllBlocks() []interface{}
	// PrintBlock just prints all the blocks.
	PrintBlock()
	// GetBlock get block at the height of this blockchain.
	GetBlock(hegiht int) (Block, error)
}

var ErrorNotFound = errors.New("block not found")
