package blockchain

import "errors"

type Block interface {
	// CalculateHash calculates hash using sha256.
	CalculateHash()
	// GetData gets data of the block.
	GetData() string
	GetPrevHash() string
}

type Blockchain interface {
	// AddBlock adds data to blockchain.
	AddBlock(data string) error
	// AllBlocks gets all the blocks of this blockchain.
	AllBlocks() ([]interface{}, error)
	// PrintBlock just prints all the blocks.
	PrintBlock()
	// GetBlock gets block at the height of this blockchain.
	GetBlock(hash string) (Block, error)
	// GetNewestHash gets newest hash of this blockchain.
	GetNewestHash() string
	// GetHeight gets height of this blockchain.
	GetHeight() int
	// Close closes blockchain.
	Close() error
}

var ErrorNotFound = errors.New("block not found")
var ErrorUnknown = errors.New("unknown")
