package bolt

import (
	"crypto/sha256"
	"fmt"
	"os"
	"testing"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/database"
	"github.com/bang9211/ossicones/mocks"
	"github.com/bang9211/ossicones/utils"
	wirejacket "github.com/bang9211/wire-jacket"

	"github.com/stretchr/testify/assert"
)

var saveblocktests = []struct {
	title    string
	input    string
	expected string
}{
	{"Saving 1 block", "Test Data1", "Test Data1"},
	{"Saving 2 block", "Test Data2", "Test Data2"},
	{"Saving 3 block", "Test Data3", "Test Data3"},
	{"Saving 4 block", "Test Data4", "Test Data4"},
}

var saveblockchaintests = []struct {
	title string
	input string
}{
	{"Saving blockchain with 1 block", "Test Data1"},
	{"Saving blockchain with 2 block", "Test Data2"},
	{"Saving blockchain with 3 block", "Test Data3"},
	{"Saving blockchain with 4 block", "Test Data4"},
}

func TestImplementBolt(t *testing.T) {
	assert.Implements(t, (*database.Database)(nil), new(BoltDB),
		"It must assert.Implements of interface database.Database")
}

func TestSaveAndGetBlock(t *testing.T) {
	cfg, db, bc, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	prevHash := ""
	height := 1
	for _, test := range saveblocktests {
		t.Run(test.title, func(t *testing.T) {
			block := &mocks.BlockMock{
				Data:     test.input,
				PrevHash: prevHash,
				Height:   height,
			}

			payload := block.Data + block.PrevHash + fmt.Sprintf("%d", block.Height)
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
			byteBlock, err := utils.ToBytes(block)
			assert.NoError(t, err)

			err = db.SaveBlock(hash, byteBlock)
			assert.NoError(t, err)

			byteBlock, err = db.GetBlock(hash)
			assert.NoError(t, err)

			newBlock := &mocks.BlockMock{}
			err = utils.FromBytes(newBlock, byteBlock)
			assert.NoError(t, err)

			assert.Equal(t, test.expected, newBlock.Data)
			assert.Equal(t, prevHash, newBlock.PrevHash)
			assert.Equal(t, height, newBlock.Height)

			prevHash = hash
			height++
		})
	}
}

func TestSaveAndGetBlockchain(t *testing.T) {
	cfg, db, bc, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	data, err := utils.ToBytes(bc)
	assert.NoError(t, err)

	err = db.SaveBlockchain(data)
	assert.NoError(t, err)

	data, err = db.GetBlockchain()
	assert.NoError(t, err)

	newBC := &mocks.BlockchainMock{}
	err = utils.FromBytes(newBC, data)
	assert.NoError(t, err)

	assert.Equal(t, bc.GetNewestHash(), newBC.NewestHash)
	assert.Equal(t, bc.GetHeight(), newBC.Height)

	for _, test := range saveblockchaintests {
		t.Run(test.title, func(t *testing.T) {
			bc.AddBlock(test.input)

			data, err = utils.ToBytes(bc)
			assert.NoError(t, err)

			err = db.SaveBlockchain(data)
			assert.NoError(t, err)

			newBC = &mocks.BlockchainMock{}
			err = utils.FromBytes(newBC, data)
			assert.NoError(t, err)

			assert.Equal(t, bc.GetNewestHash(), newBC.NewestHash)
			assert.Equal(t, bc.GetHeight(), newBC.Height)
		})
	}
}

func initTest() (config.Config, database.Database, blockchain.Blockchain, error) {
	cfg := wirejacket.GetConfig()

	os.Remove("ossicones.db")

	db := New(cfg)
	bc := &mocks.BlockchainMock{}
	bc.Init()

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
