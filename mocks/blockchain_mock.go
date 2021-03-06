package mocks

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/stretchr/testify/mock"
)

const defaultDifficulty = 2

type BlockchainMock struct {
	mock.Mock
	NewestHash string `json:"newstHash"`
	Height     int    `json:"height"`
	blocks     []blockchain.Block
	blocksMap  map[string]blockchain.Block
}

func (b *BlockchainMock) Init() {
	b.blocksMap = make(map[string]blockchain.Block)

	b.blocks = make([]blockchain.Block, 0)
	b.AddBlock(os.Getenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA"))
}

func (b *BlockchainMock) AddBlock(data string) error {
	newBlock := &BlockMock{
		Data:     data,
		PrevHash: b.NewestHash,
		Height:   b.Height + 1,
	}
	payload := newBlock.Data + newBlock.PrevHash + fmt.Sprintf("%d", newBlock.Height)
	newBlock.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))

	b.blocks = append(b.blocks, newBlock)
	b.blocksMap[newBlock.Hash] = newBlock

	return nil
}

func (b *BlockchainMock) AllBlocks() ([]blockchain.Block, error) {
	return b.blocks, nil
}

func (b *BlockchainMock) PrintBlock() {}

func (b *BlockchainMock) GetBlock(hash string) (blockchain.Block, error) {
	return b.blocksMap[hash], nil
}

func (b *BlockchainMock) GetNewestHash() string { return b.NewestHash }

func (b *BlockchainMock) GetHeight() int { return b.Height }

func (b *BlockchainMock) Close() error { return nil }

type BlockMock struct {
	mock.Mock
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevhash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

func (b *BlockMock) Mine() {
	target := strings.Repeat("0", defaultDifficulty)
	for {
		byteBlock := []byte(fmt.Sprint(b))
		hash := fmt.Sprintf("%x", sha256.Sum256(byteBlock))
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			return
		} else {
			b.Nonce++
		}
	}
}

func (b *BlockMock) GetData() string {
	return b.Data
}

func (b *BlockMock) GetPrevHash() string {
	return b.PrevHash
}
