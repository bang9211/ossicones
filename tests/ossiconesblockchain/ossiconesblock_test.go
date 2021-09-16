package ossiconesblockchain

import (
	"testing"

	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/blockchain"

	. "github.com/stretchr/testify/assert"
)

var calculatehashtests = []struct {
	title    string
	input    string
	expected string
}{
	{"Calculating hash of Test Data1", "Test Data1", "bfa6723ef7f9771b84a4f6cc43d1d5e80c1bbcf965def63f7af70c392c4839e2"},
	{"Calculating hash of Test Data2", "Test Data2", "0db581d518bb6f6d995afd26e6d3f35aebd207b8a1e4d5a3851ccdca1be88209"},
	{"Calculating hash of Test Data3", "Test Data3", "84084c8d4b8b7694b15ec1d5c6ce23a2e68704be54508b7817cabaf76503510f"},
	{"Calculating hash of Test Data4", "Test Data4", "19476fec339cc95ca6b8325eceb99f944619d69fd9c37f7fc2fa536094750538"},
}

var getdatatests = []struct {
	title    string
	input    string
	expected string
}{
	{"Getting data of block1", "Test Data1", "Test Data1"},
	{"Getting data of block2", "Test Data2", "Test Data2"},
	{"Getting data of block3", "Test Data3", "Test Data3"},
	{"Getting data of block4", "Test Data4", "Test Data4"},
}

func TestImplementsBlock(t *testing.T) {
	Implements(t, (*blockchain.Block)(nil), new(ossiconesblockchain.OssiconesBlock),
		"It must implements of interface blockchain.Block")
}

func TestCalculateHash(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer NoError(t, closeTest(cfg, bc), "Failed to closeTest()")

	blocks := bc.AllBlocks()
	Equal(t, 1, len(blocks))

	Implements(t, (*blockchain.Block)(nil), blocks[len(blocks)-1],
		"It must implements of interface blockchain.Block")
	block, ok := blocks[len(blocks)-1].(blockchain.Block)
	True(t, ok)

	IsType(t, (*ossiconesblockchain.OssiconesBlock)(nil), block,
		"It should be equal of type ossiconesblockchain.OssiconesBlock")
	ossiconesBlock, ok := block.(*ossiconesblockchain.OssiconesBlock)
	True(t, ok)

	Equal(t, ossiconesBlock.Hash,
		"90aa8185295e7f87c0c1608967e4abf6a5e201938180a3a6cc8d891d51283532")

	for _, test := range calculatehashtests {
		t.Run(test.title, func(t *testing.T) {
			newBlock := ossiconesblockchain.OssiconesBlock{
				Data:     test.input,
				PrevHash: ossiconesBlock.Hash,
				Height:   len(blocks) + 1,
			}
			newBlock.CalculateHash()
			Equal(t, test.expected, newBlock.Hash)

			bc.AddBlock(test.input)
			blocks = bc.AllBlocks()
			Implements(t, (*blockchain.Block)(nil), blocks[len(blocks)-1],
				"It must implements of interface blockchain.Block")
			block, ok := blocks[len(blocks)-1].(blockchain.Block)
			True(t, ok)

			IsType(t, (*ossiconesblockchain.OssiconesBlock)(nil), block,
				"It should be equal of type ossiconesblockchain.OssiconesBlock")
			ossiconesBlock, ok = block.(*ossiconesblockchain.OssiconesBlock)
			True(t, ok)
		})
	}
}

func TestGetData(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer NoError(t, closeTest(cfg, bc), "Failed to closeTest()")

	blocks := bc.AllBlocks()
	Equal(t, 1, len(blocks))

	Implements(t, (*blockchain.Block)(nil), blocks[len(blocks)-1],
		"It must implements of interface blockchain.Block")
	block, ok := blocks[len(blocks)-1].(blockchain.Block)
	True(t, ok)

	Equal(t, block.GetData(), genesisBlockData)

	for _, test := range getdatatests {
		t.Run(test.title, func(t *testing.T) {
			bc.AddBlock(test.input)
			blocks = bc.AllBlocks()
			Implements(t, (*blockchain.Block)(nil), blocks[len(blocks)-1],
				"It must implements of interface blockchain.Block")
			block, ok := blocks[len(blocks)-1].(blockchain.Block)
			True(t, ok)
			Equal(t, test.expected, block.GetData())
		})
	}
}
