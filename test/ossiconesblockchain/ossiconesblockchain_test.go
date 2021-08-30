package ossiconesblockchain

import (
	"testing"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/modules"

	. "github.com/stretchr/testify/assert"
)

var addblocktests = []struct {
	title string
	in    string
	out   int
}{
	{"Added 1 block", "Test Data1", 2},
	{"Added 2 block", "Test Data2", 3},
	{"Added 3 block", "Test Data3", 4},
	{"Added 4 block", "Test Data4", 5},
}

func TestAddBlock(t *testing.T) {
	config, err := modules.InitConfig()
	Nil(t, err)

	bc, err := modules.InitBlockchain(config)
	Nil(t, err)

	blocks := bc.AllBlocks()
	Equal(t, len(blocks), 1)

	block, success := blocks[0].(blockchain.Block)
	True(t, success)
	Equal(t, "Genesis OssiconesBlock", block.GetData())

	for i, test := range addblocktests {
		t.Run(test.title, func(t *testing.T) {
			bc.AddBlock(test.in)

			blocks = bc.AllBlocks()
			Equal(t, len(blocks), test.out)

			block, success := blocks[i+1].(blockchain.Block)
			True(t, success)
			Equal(t, test.in, block.GetData())
		})
	}
}

func TestAllBlocks(t *testing.T) {

}

func TestPrintBlock(t *testing.T) {

}

func TestGetBlock(t *testing.T) {

}
