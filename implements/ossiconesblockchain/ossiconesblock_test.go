package ossiconesblockchain

import (
	"testing"

	"github.com/bang9211/ossicones/interfaces/blockchain"

	"github.com/stretchr/testify/assert"
)

var calculatehashtests = []struct {
	title    string
	input    string
	expected string
}{
	{"Calculating hash of Test Data1", "Test Data1", "008059a4ca5ace8c990bd225e797bcac6da370efcd95d3b3f858188055c51a33"},
	{"Calculating hash of Test Data2", "Test Data2", "00421e492d1fba9d65a13a48fd28c7884aabf0563c0795f0e87a59df2e6a50ea"},
	{"Calculating hash of Test Data3", "Test Data3", "004dec6a66243fc6f97f89b762653b77a3ccde9bde8f1e15aaa840f92ef086a1"},
	{"Calculating hash of Test Data4", "Test Data4", "00cc036c41548658773c19a0bc0f0abcb08f4e50b4e549b339663b395a986765"},
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

var getprevhashtests = []struct {
	title    string
	input    string
	expected string
}{
	{"Getting Previous hash of block1", "Test Data1", "00a283a808beba990e639df72bdc43dea322513dbd56ec55257a3ed7e500ff8b"},
	{"Getting Previous hash of block2", "Test Data2", "00ef0a9819b1685d07b18d9a17a9b34061fd024cb2c9c1e7adb1b56f2233fe4c"},
	{"Getting Previous hash of block3", "Test Data3", "00e679f23b50b25daa268358bd6c7a8a78692cdd1c9affc7e7fb906cb18378aa"},
	{"Getting Previous hash of block4", "Test Data4", "00bc734cbcb4ee90e0579364f1203868556347fefee030f791f89757774b65b0"},
}

func TestImplementsBlock(t *testing.T) {
	assert.Implements(t, (*blockchain.Block)(nil), new(OssiconesBlock),
		"It must implements of interface blockchain.Block")
}

func TestMine(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	blocks, err := bc.AllBlocks()
	assert.Equal(t, 1, len(blocks))
	assert.NoError(t, err)

	block := blocks[len(blocks)-1]
	assert.IsType(t, (*OssiconesBlock)(nil), block,
		"It should be equal of type OssiconesBlock")

	ossiconesBlock, ok := block.(*OssiconesBlock)
	assert.True(t, ok)

	assert.Equal(t, ossiconesBlock.Hash,
		"00a283a808beba990e639df72bdc43dea322513dbd56ec55257a3ed7e500ff8b")

	for _, test := range calculatehashtests {
		t.Run(test.title, func(t *testing.T) {
			newBlock := OssiconesBlock{
				Data:     test.input,
				PrevHash: ossiconesBlock.Hash,
				Height:   len(blocks) + 1,
			}
			newBlock.Mine()
			assert.Equal(t, test.expected, newBlock.Hash)

			bc.AddBlock(test.input)
			blocks, err = bc.AllBlocks()
			assert.NoError(t, err)
			assert.Implements(t, (*blockchain.Block)(nil), blocks[len(blocks)-1],
				"It must implements of interface blockchain.Block")
			block, ok := blocks[len(blocks)-1].(blockchain.Block)
			assert.True(t, ok)

			assert.IsType(t, (*OssiconesBlock)(nil), block,
				"It should be equal of type ossiconesblockchain.OssiconesBlock")
			ossiconesBlock, ok = block.(*OssiconesBlock)
			assert.True(t, ok)
		})
	}
}

func TestGetData(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	blocks, err := bc.AllBlocks()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(blocks))

	assert.Implements(t, (*blockchain.Block)(nil), blocks[len(blocks)-1],
		"It must implements of interface blockchain.Block")

	assert.Equal(t, blocks[0].GetData(), genesisBlockData)

	for _, test := range getdatatests {
		t.Run(test.title, func(t *testing.T) {
			err := bc.AddBlock(test.input)
			assert.NoError(t, err)
			hash := bc.GetNewestHash()
			block, err := bc.GetBlock(hash)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, block.GetData())
		})
	}
}

func TestGetPrevHash(t *testing.T) {
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, db, bc, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db, bc)

	blocks, err := bc.AllBlocks()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(blocks))

	assert.Equal(t, blocks[0].GetData(), genesisBlockData)
	for _, test := range getprevhashtests {
		t.Run(test.title, func(t *testing.T) {
			err := bc.AddBlock(test.input)
			assert.NoError(t, err)
			hash := bc.GetNewestHash()
			block, err := bc.GetBlock(hash)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, block.GetPrevHash())
		})
	}
}
