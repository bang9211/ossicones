package defaultexplorerserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bang9211/ossicones/implements/bolt"
	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/explorerserver"
	"github.com/bang9211/ossicones/utils"
	wirejacket "github.com/bang9211/wire-jacket"

	"github.com/stretchr/testify/assert"
)

const genesisBlockData = "TEST_GENESIS_BLOCK_DATA"

func TestImplementExplorerServer(t *testing.T) {
	assert.Implements(t, (*explorerserver.ExplorerServer)(nil), new(DefaultExplorerServer),
		"It must assert.Implements of interface explorerserver.ExplorerServer")
}

func TestServe(t *testing.T) {
	_, err := utils.GetOrSetHomePath()
	assert.NoError(t, err)
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, es, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, bc, es)

	des, ok := es.(*DefaultExplorerServer)
	assert.True(t, ok)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	responseRecorder := httptest.NewRecorder()
	des.home(responseRecorder, request)
	assert.Equal(t, responseRecorder.Code, 200)
}

func TestClose(t *testing.T) {
	assert.Implements(t, (*explorerserver.ExplorerServer)(nil), new(DefaultExplorerServer),
		"It must assert.Implements of interface explorerserver.ExplorerServer")
}

func initTest() (config.Config, blockchain.Blockchain, explorerserver.ExplorerServer, error) {
	cfg := wirejacket.GetConfig()

	db := bolt.New(cfg)
	bc := ossiconesblockchain.New(cfg, db)
	es := New(cfg, bc)

	return cfg, bc, es, nil
}

func closeTest(cfg config.Config, bc blockchain.Blockchain, es explorerserver.ExplorerServer) error {
	err := es.Close()
	if err != nil {
		return err
	}
	err = bc.Close()
	if err != nil {
		return err
	}
	return cfg.Close()
}
