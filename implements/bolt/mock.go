package bolt

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/database"
	wirejacket "github.com/bang9211/wire-jacket"
	"github.com/stretchr/testify/mock"
)

const genesisBlockData = "TEST_GENESIS_BLOCK_DATA"

type blockchainMock struct {
	mock.Mock
	NewestHash string `json:"newstHash"`
	Height     int    `json:"height"`
	blocks     []blockchain.Block
	blocksMap  map[string]blockchain.Block
}

func (b *blockchainMock) AddBlock(data string) error {
	newBlock := &blockMock{
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

func (b *blockchainMock) AllBlocks() ([]blockchain.Block, error) {
	return b.blocks, nil
}

func (b *blockchainMock) PrintBlock() {}

func (b *blockchainMock) GetBlock(hash string) (blockchain.Block, error) {
	return b.blocksMap[hash], nil
}

func (b *blockchainMock) GetNewestHash() string { return b.NewestHash }

func (b *blockchainMock) GetHeight() int { return b.Height }

func (b *blockchainMock) Close() error { return nil }

type blockMock struct {
	mock.Mock
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevhash,omitempty"`
	Height   int    `json:"height"`
}

func (b *blockMock) CalculateHash() {
	Hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", Hash)
}

func (b *blockMock) GetData() string {
	return b.Data
}

func (b *blockMock) GetPrevHash() string {
	return b.PrevHash
}

func initTest() (config.Config, database.Database, blockchain.Blockchain, error) {
	cfg := wirejacket.GetConfig()

	os.Remove("ossicones.db")

	db := New(cfg)
	bc := &blockchainMock{
		blocks:    []blockchain.Block{},
		blocksMap: map[string]blockchain.Block{},
	}

	return cfg, db, bc, nil
}

func closeTest(cfg config.Config, db database.Database, bc blockchain.Blockchain) error {
	err := bc.Close()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}

	os.Remove("ossicones.db")

	return cfg.Close()
}
