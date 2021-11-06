package bolt

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/bang9211/ossicones/implements/bolt"
	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/database"
	"github.com/bang9211/ossicones/utils"

	. "github.com/stretchr/testify/assert"
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

func TestImplementBolt(t *testing.T) {
	Implements(t, (*database.Database)(nil), new(bolt.BoltDB),
		"It must implements of interface database.Database")
}

func TestSaveBlock(t *testing.T) {
	cfg, db, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, db)

	prevHash := ""
	height := 1
	for _, test := range saveblocktests {
		t.Run(test.title, func(t *testing.T) {
			block := &ossiconesblockchain.OssiconesBlock{
				Data:     test.input,
				PrevHash: prevHash,
				Height:   height,
			}

			payload := block.Data + block.PrevHash + fmt.Sprintf("%d", block.Height)
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
			byteBlock, err := utils.ToBytes(block)
			NoError(t, err)

			err = db.SaveBlock(hash, byteBlock)
			NoError(t, err)

			byteBlock, err = db.GetBlock(hash)
			NoError(t, err)

			newBlock := &ossiconesblockchain.OssiconesBlock{}
			err = utils.FromBytes(newBlock, byteBlock)
			NoError(t, err)

			Equal(t, test.expected, newBlock.Data)
			Equal(t, prevHash, newBlock.PrevHash)
			Equal(t, height, newBlock.Height)

			prevHash = hash
			height++
		})
	}

}

func TestSaveBlockchain(t *testing.T) {
	Implements(t, (*database.Database)(nil), new(bolt.BoltDB),
		"It must implements of interface database.Database")
}

func TestGetCheckpoint(t *testing.T) {
	Implements(t, (*database.Database)(nil), new(bolt.BoltDB),
		"It must implements of interface database.Database")
}

func TestGetBlock(t *testing.T) {
	Implements(t, (*database.Database)(nil), new(bolt.BoltDB),
		"It must implements of interface database.Database")
}

func TestClose(t *testing.T) {
	Implements(t, (*database.Database)(nil), new(bolt.BoltDB),
		"It must implements of interface database.Database")
}
