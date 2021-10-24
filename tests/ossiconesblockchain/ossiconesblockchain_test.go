package ossiconesblockchain

import (
	"testing"

	"github.com/bang9211/ossicones/interfaces/blockchain"

	. "github.com/stretchr/testify/assert"
)

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

func TestImplementsBlockchain(t *testing.T) {
	cfg, db, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer NoError(t, closeTest(cfg, db, bc), "Failed to closeTest()")

	Implements(t, (*blockchain.Blockchain)(nil), bc,
		"It must implements of interface blockchain.Blockchain")
}

func TestAddBlock(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer NoError(t, closeTest(cfg, db, bc), "Failed to closeTest()")

	blocks, err := bc.AllBlocks()
	NoError(t, err)
	Equal(t, 1, len(blocks))

	Implements(t, (*blockchain.Block)(nil), blocks[0],
		"It must implements of interface blockchain.Block")
	block, ok := blocks[0].(blockchain.Block)
	True(t, ok)
	Equal(t, genesisBlockData, block.GetData())

	for i, test := range addblocktests {
		t.Run(test.title, func(t *testing.T) {
			bc.AddBlock(test.input)

			blocks, err = bc.AllBlocks()
			NoError(t, err)
			Equal(t, test.expected, len(blocks))

			Implements(t, (*blockchain.Block)(nil), blocks[i+1],
				"It must implements of interface blockchain.Block")
			block, ok := blocks[i+1].(blockchain.Block)
			True(t, ok)
			Equal(t, test.input, block.GetData())
		})
	}
}

func TestAllBlocks(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer NoError(t, closeTest(cfg, db, bc), "Failed to closeTest()")

	blocks, err := bc.AllBlocks()
	NoError(t, err)
	Equal(t, 1, len(blocks))

	Implements(t, (*blockchain.Block)(nil), blocks[0],
		"It must implements of interface blockchain.Block")
	block, ok := blocks[0].(blockchain.Block)
	True(t, ok)
	Equal(t, genesisBlockData, block.GetData())

	for i, test := range allblockstests {
		t.Run(test.title, func(t *testing.T) {
			bc.AddBlock(test.input)

			blocks, err = bc.AllBlocks()
			NoError(t, err)
			Equal(t, test.expected, len(blocks))

			Implements(t, (*blockchain.Block)(nil), blocks[i+1],
				"It must implements of interface blockchain.Block")
			block, ok := blocks[i+1].(blockchain.Block)
			True(t, ok)
			Equal(t, test.input, block.GetData())
		})
	}
}

func TestGetBlock(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer NoError(t, closeTest(cfg, db, bc), "Failed to closeTest()")

	blocks, err := bc.AllBlocks()
	NoError(t, err)
	Equal(t, 1, len(blocks))

	Implements(t, (*blockchain.Block)(nil), blocks[0],
		"It must implements of interface blockchain.Block")
	block, ok := blocks[0].(blockchain.Block)
	True(t, ok)
	Equal(t, genesisBlockData, block.GetData())

	//todo get genesisblock
	genesisBlock, err := bc.GetBlock(bc.GetNewestHash())
	NoError(t, err, "Failed to get a generation block")
	Equal(t, genesisBlockData, genesisBlock.GetData())

	blockCount := bc.GetHeight()
	for _, test := range getblocktests {
		t.Run(test.title, func(t *testing.T) {
			err := bc.AddBlock(test.input_data)
			NoError(t, err)
			block, err := bc.GetBlock(bc.GetNewestHash())
			NoError(t, err, "Failed to get a last block")
			Equal(t, test.input_data, block.GetData())
			blockCount++
		})
	}

	blocks, err = bc.AllBlocks()
	NoError(t, err)
	Equal(t, len(blocks), blockCount)

	hash := bc.GetNewestHash()
	for _, test := range getblocktests {
		t.Run(test.title, func(t *testing.T) {
			block, err := bc.GetBlock(hash)
			NoError(t, err, "Failed to get a block of index(%d)", test.input_data)
			Equal(t, test.input_data, block.GetData())
			hash = block.GetPrevHash()
		})
	}
	blocks, err = bc.AllBlocks()
	NoError(t, err)
	Equal(t, len(blocks), blockCount)
}
