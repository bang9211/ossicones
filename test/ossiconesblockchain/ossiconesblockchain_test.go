package ossiconesblockchain

import (
	"testing"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/modules"

	. "github.com/stretchr/testify/assert"
)

var addblocktests = []struct {
	title    string
	input    string
	expected int
}{
	{"Added 1 block", "Test Data1", 2},
	{"Added 2 block", "Test Data2", 3},
	{"Added 3 block", "Test Data3", 4},
	{"Added 4 block", "Test Data4", 5},
}

var allblockstests = []struct {
	title    string
	input    string
	expected int
}{
	{"Existed 2 block", "Test Data1", 2},
	{"Existed 3 block", "Test Data2", 3},
	{"Existed 4 block", "Test Data3", 4},
	{"Existed 5 block", "Test Data4", 5},
}

var getblocktests = []struct {
	title       string
	input_data  string
	input_index int
	expected    string
}{
	{"Get Test Data1 block", "Test Data1", 2, "Test Data1"},
	{"Get Test Data2 block", "Test Data2", 3, "Test Data2"},
	{"Get Test Data3 block", "Test Data3", 4, "Test Data3"},
	{"Get Test Data4 block", "Test Data4", 5, "Test Data4"},
}

const genesisBlockData = "TEST_GENESIS_BLOCK_DATA"

func TestAddBlock(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, bc)

	blocks := bc.AllBlocks()
	Equal(t, 1, len(blocks))

	Implements(t, (*blockchain.Block)(nil), blocks[0],
		"It must implements of interface blockchain.Block")
	block, success := blocks[0].(blockchain.Block)
	True(t, success)
	Equal(t, genesisBlockData, block.GetData())

	for i, test := range addblocktests {
		t.Run(test.title, func(t *testing.T) {
			bc.AddBlock(test.input)

			blocks = bc.AllBlocks()
			Equal(t, test.expected, len(blocks))

			Implements(t, (*blockchain.Block)(nil), blocks[i+1],
				"It must implements of interface blockchain.Block")
			block, success := blocks[i+1].(blockchain.Block)
			True(t, success)
			Equal(t, test.input, block.GetData())
		})
	}
}

func TestAllBlocks(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, bc)

	blocks := bc.AllBlocks()
	Equal(t, 1, len(blocks))

	Implements(t, (*blockchain.Block)(nil), blocks[0],
		"It must implements of interface blockchain.Block")
	block, success := blocks[0].(blockchain.Block)
	True(t, success)
	Equal(t, genesisBlockData, block.GetData())

	for i, test := range allblockstests {
		t.Run(test.title, func(t *testing.T) {
			bc.AddBlock(test.input)

			blocks = bc.AllBlocks()
			Equal(t, test.expected, len(blocks))

			Implements(t, (*blockchain.Block)(nil), blocks[i+1],
				"It must implements of interface blockchain.Block")
			block, success := blocks[i+1].(blockchain.Block)
			True(t, success)
			Equal(t, test.input, block.GetData())
		})
	}
}

func TestGetBlock(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, bc)

	blocks := bc.AllBlocks()
	Equal(t, 1, len(blocks))

	Implements(t, (*blockchain.Block)(nil), blocks[0],
		"It must implements of interface blockchain.Block")
	block, success := blocks[0].(blockchain.Block)
	True(t, success)
	Equal(t, genesisBlockData, block.GetData())

	genesisBlock, err := bc.GetBlock(1)
	NoError(t, err, "Failed to get a generation block")
	Equal(t, genesisBlockData, genesisBlock.GetData())

	blockCount := len(bc.AllBlocks())
	for _, test := range getblocktests {
		t.Run(test.title, func(t *testing.T) {
			bc.AddBlock(test.input_data)

			block, err := bc.GetBlock(len(bc.AllBlocks()))
			NoError(t, err, "Failed to get a last block")
			Equal(t, test.input_data, block.GetData())
			blockCount++
		})
	}

	Equal(t, len(bc.AllBlocks()), blockCount)

	for _, test := range getblocktests {
		t.Run(test.title, func(t *testing.T) {
			block, err := bc.GetBlock(test.input_index)
			NoError(t, err, "Failed to get a block of index(%d)", test.input_data)
			Equal(t, test.input_data, block.GetData())
		})
	}
	_, err = bc.GetBlock(len(bc.AllBlocks()) + 1)
	ErrorIs(t, err, blockchain.ErrorNotFound, "There is no Block, the error must be blockchain.ErrorNotFound")
}

func TestReset(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, bc)

	blocks := bc.AllBlocks()
	Equal(t, 1, len(blocks))

	for i := 0; i < 5; i++ {
		bc.AddBlock("Test Data")
	}
	blocks = bc.AllBlocks()
	Equal(t, 6, len(blocks))

	NoError(t, bc.Reset(), "Reset must be successed")

	blocks = bc.AllBlocks()
	Equal(t, 1, len(blocks))
}

//init
func initTest() (config.Config, blockchain.Blockchain, error) {
	cfg, err := modules.InitConfig()
	if err != nil {
		return nil, nil, err
	}

	bc, err := modules.InitBlockchain(cfg)
	if err != nil {
		return nil, nil, err
	}

	err = bc.Reset()
	if err != nil {
		return nil, nil, err
	}

	return cfg, bc, nil
}

//close
func closeTest(cfg config.Config, bc blockchain.Blockchain) error {
	err := bc.Close()
	if err != nil {
		return err
	}
	return cfg.Close()
}
