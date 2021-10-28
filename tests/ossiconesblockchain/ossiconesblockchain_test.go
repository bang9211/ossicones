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
	defer closeTest(cfg, db, bc)

	Implements(t, (*blockchain.Blockchain)(nil), bc,
		"It must implements of interface blockchain.Blockchain")
}

func TestAddBlock(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	blocks, err := bc.AllBlocks()
	NoError(t, err)
	Equal(t, 1, len(blocks))

	Equal(t, genesisBlockData, blocks[0].GetData())

	for _, test := range addblocktests {
		t.Run(test.title, func(t *testing.T) {
			bc.AddBlock(test.input)
			block, err := bc.GetBlock(bc.GetNewestHash())
			NoError(t, err)
			Equal(t, test.input, block.GetData())
		})
	}
}

func TestAllBlocks(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	blocks, err := bc.AllBlocks()
	NoError(t, err)
	Equal(t, 1, len(blocks))

	Equal(t, genesisBlockData, blocks[0].GetData())

	for _, test := range allblockstests {
		t.Run(test.title, func(t *testing.T) {
			bc.AddBlock(test.input)

			blocks, err = bc.AllBlocks()
			NoError(t, err)
			Equal(t, test.expected, len(blocks))
			Equal(t, test.input, blocks[0].GetData())
		})
	}
}

func TestGetBlock(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	blocks, err := bc.AllBlocks()
	NoError(t, err)
	Equal(t, 1, len(blocks))

	Equal(t, genesisBlockData, blocks[0].GetData())

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
			NoError(t, err)
			Equal(t, test.expected, block.GetData())
			hash = block.GetPrevHash()
		})
	}
	blocks, err = bc.AllBlocks()
	NoError(t, err)
	Equal(t, len(blocks), blockCount)
}
