package ossiconesblockchain

import (
	"fmt"
	"os"
	"testing"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/database"
	"github.com/bang9211/ossicones/mocks"
	wirejacket "github.com/bang9211/wire-jacket"

	"github.com/stretchr/testify/assert"
)

const genesisBlockData = "TEST_GENESIS_BLOCK_DATA"

var addblocktests = []struct {
	title    string
	input    string
	expected int
}{
	{"Adding 1 block", "Test Data1", 2},
	{"Adding 2 block", "Test Data2", 3},
	{"Adding 3 block", "Test Data3", 4},
	{"Adding 4 block", "Test Data4", 5},
}

var allblockstests = []struct {
	title    string
	input    string
	expected int
}{
	{"Checking 2 block", "Test Data1", 2},
	{"Checking 3 block", "Test Data2", 3},
	{"Checking 4 block", "Test Data3", 4},
	{"Checking 5 block", "Test Data4", 5},
}

var getblocktests = []struct {
	title      string
	input_data string
	expected   string
}{
	{"Getting Test Data1 block", "Test Data1", "Test Data4"},
	{"Getting Test Data2 block", "Test Data2", "Test Data3"},
	{"Getting Test Data3 block", "Test Data3", "Test Data2"},
	{"Getting Test Data4 block", "Test Data4", "Test Data1"},
}

var getnewesthashtests = []struct {
	title      string
	input_data string
	expected   string
}{
	{"Getting Test Data1 block", "Test Data1", "a9cbc6f70a1af8ffc003e3a1a9ef87d41f4b3113c66c1b2663625601609012f1"},
	{"Getting Test Data2 block", "Test Data2", "7298f26aa20f68ec9c2fb751d6e8168f25300632cf904c9db0fd1acb42b61eec"},
	{"Getting Test Data3 block", "Test Data3", "4bf189892f3dd47db879e79e2e604eaefb76831a2b9fb87ccb6ede5e93aad126"},
	{"Getting Test Data4 block", "Test Data4", "fe5bb6f93039110a918aaccbf01fd8ee399acb32925931bab3fb4881d2ebef81"},
}

var getheighttests = []struct {
	title      string
	input_data string
	expected   int
}{
	{"Getting Test Data1 block", "Test Data1", 2},
	{"Getting Test Data2 block", "Test Data2", 3},
	{"Getting Test Data3 block", "Test Data3", 4},
	{"Getting Test Data4 block", "Test Data4", 5},
}

func TestImplementsBlockchain(t *testing.T) {
	assert.Implements(t, (*blockchain.Blockchain)(nil), new(OssiconesBlockchain),
		"It must implements of interface blockchain.Blockchain")
}

func TestAddBlock(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	blocks, err := bc.AllBlocks()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(blocks))

	assert.Equal(t, genesisBlockData, blocks[0].GetData())

	for _, test := range addblocktests {
		t.Run(test.title, func(t *testing.T) {
			bc.AddBlock(test.input)
			block, err := bc.GetBlock(bc.GetNewestHash())
			assert.NoError(t, err)
			assert.Equal(t, test.input, block.GetData())
		})
	}
}

func TestAllBlocks(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	blocks, err := bc.AllBlocks()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(blocks))

	assert.Equal(t, genesisBlockData, blocks[0].GetData())

	for _, test := range allblockstests {
		t.Run(test.title, func(t *testing.T) {
			bc.AddBlock(test.input)

			blocks, err = bc.AllBlocks()
			assert.NoError(t, err)
			assert.Equal(t, test.expected, len(blocks))
			assert.Equal(t, test.input, blocks[0].GetData())
		})
	}
}

func TestGetBlock(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	blocks, err := bc.AllBlocks()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(blocks))
	assert.Equal(t, genesisBlockData, blocks[0].GetData())

	genesisBlock, err := bc.GetBlock(bc.GetNewestHash())
	assert.NoError(t, err, "Failed to get a generation block")
	assert.Equal(t, genesisBlockData, genesisBlock.GetData())

	blockCount := bc.GetHeight()
	for _, test := range getblocktests {
		t.Run(test.title, func(t *testing.T) {
			err := bc.AddBlock(test.input_data)
			assert.NoError(t, err)
			block, err := bc.GetBlock(bc.GetNewestHash())
			assert.NoError(t, err, "Failed to get a last block")
			assert.Equal(t, test.input_data, block.GetData())
			blockCount++
		})
	}

	blocks, err = bc.AllBlocks()
	assert.NoError(t, err)
	assert.Equal(t, len(blocks), blockCount)

	hash := bc.GetNewestHash()
	for _, test := range getblocktests {
		t.Run(test.title, func(t *testing.T) {
			block, err := bc.GetBlock(hash)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, block.GetData())
			hash = block.GetPrevHash()
		})
	}
	blocks, err = bc.AllBlocks()
	assert.NoError(t, err)
	assert.Equal(t, len(blocks), blockCount)
}

func TestGetNewestHash(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	blocks, err := bc.AllBlocks()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(blocks))
	assert.Equal(t, genesisBlockData, blocks[0].GetData())

	genesisBlock, err := bc.GetBlock(bc.GetNewestHash())
	assert.NoError(t, err, "Failed to get a generation block")
	assert.Equal(t, genesisBlockData, genesisBlock.GetData())

	for _, test := range getnewesthashtests {
		t.Run(test.title, func(t *testing.T) {
			err := bc.AddBlock(test.input_data)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, bc.GetNewestHash())
		})
	}
}

func TestGetHeight(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	blocks, err := bc.AllBlocks()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(blocks))
	assert.Equal(t, genesisBlockData, blocks[0].GetData())

	genesisBlock, err := bc.GetBlock(bc.GetNewestHash())
	assert.NoError(t, err, "Failed to get a generation block")
	assert.Equal(t, genesisBlockData, genesisBlock.GetData())

	height := bc.GetHeight()
	assert.Equal(t, height, 1)
	for _, test := range getheighttests {
		t.Run(test.title, func(t *testing.T) {
			err := bc.AddBlock(test.input_data)
			assert.NoError(t, err)
			_, err = bc.GetBlock(bc.GetNewestHash())
			assert.NoError(t, err, "Failed to get a last block")
			assert.Equal(t, test.expected, bc.GetHeight())
		})
	}
}

func initTest() (config.Config, database.Database, blockchain.Blockchain, error) {
	cfg := wirejacket.GetConfig()

	os.Remove("test.db")

	bc := New(cfg, &mocks.DBMock{Blocks: map[string][]byte{}})
	if bc == nil {
		return nil, nil, nil, fmt.Errorf("failed to New()")
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
