package ossiconesblockchain

import (
	"os"
	"testing"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/database"
	wirejacket "github.com/bang9211/wire-jacket"

	"github.com/stretchr/testify/assert"
)

var calculatehashtests = []struct {
	title    string
	input    string
	expected string
}{
	{"Calculating hash of Test Data1", "Test Data1", "ed739e9aeb8d09d33a8687fa6c35d88b5bcf22c5ba134b7862d8347e55016262"},
	{"Calculating hash of Test Data2", "Test Data2", "f781594d46943ee46a392363426db980f8a87466d9ab53ded834c92069890330"},
	{"Calculating hash of Test Data3", "Test Data3", "54d7f6fc926e7d51a399a45a82f1af01c9bac63eb135d1815cc309a05831944e"},
	{"Calculating hash of Test Data4", "Test Data4", "e91e533a9c1e791d73c547a32240a5f7454ad7e3b205070f89861cff2ff67bef"},
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
	{"Getting Previous hash of block1", "Test Data1", "46a823ac625c5d086378bd28d032ffd421738bdb1f13f8a403486bc7ea645082"},
	{"Getting Previous hash of block2", "Test Data2", "a9cbc6f70a1af8ffc003e3a1a9ef87d41f4b3113c66c1b2663625601609012f1"},
	{"Getting Previous hash of block3", "Test Data3", "7298f26aa20f68ec9c2fb751d6e8168f25300632cf904c9db0fd1acb42b61eec"},
	{"Getting Previous hash of block4", "Test Data4", "4bf189892f3dd47db879e79e2e604eaefb76831a2b9fb87ccb6ede5e93aad126"},
}

func TestImplementsBlock(t *testing.T) {
	assert.Implements(t, (*blockchain.Block)(nil), new(OssiconesBlock),
		"It must implements of interface blockchain.Block")
}

func TestCalculateHash(t *testing.T) {
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
		"46a823ac625c5d086378bd28d032ffd421738bdb1f13f8a403486bc7ea645082")

	for _, test := range calculatehashtests {
		t.Run(test.title, func(t *testing.T) {
			newBlock := OssiconesBlock{
				Data:     test.input,
				PrevHash: ossiconesBlock.Hash,
				Height:   len(blocks) + 1,
			}
			newBlock.CalculateHash()
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
