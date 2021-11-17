package defaultexplorerserver

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/explorerserver"
	"github.com/bang9211/ossicones/mocks"
	"github.com/bang9211/ossicones/utils"
	viperjacket "github.com/bang9211/viper-jacket"
	wirejacket "github.com/bang9211/wire-jacket"

	"github.com/stretchr/testify/assert"
)

const genesisBlockData = "TEST_GENESIS_BLOCK_DATA"

func TestImplementExplorerServer(t *testing.T) {
	assert.Implements(t, (*explorerserver.ExplorerServer)(nil), new(DefaultExplorerServer),
		"It must assert.Implements of interface explorerserver.ExplorerServer")
}

func TestGetHome(t *testing.T) {
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

func TestGetAdd(t *testing.T) {
	_, err := utils.GetOrSetHomePath()
	assert.NoError(t, err)
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, es, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, bc, es)

	des, ok := es.(*DefaultExplorerServer)
	assert.True(t, ok)

	request := httptest.NewRequest(http.MethodGet, "/add", nil)
	responseRecorder := httptest.NewRecorder()
	des.add(responseRecorder, request)
	assert.Equal(t, responseRecorder.Code, 200)
}

func TestPostAdd(t *testing.T) {
	_, err := utils.GetOrSetHomePath()
	assert.NoError(t, err)
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, es, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, bc, es)

	des, ok := es.(*DefaultExplorerServer)
	assert.True(t, ok)

	data := url.Values{}
	data.Set("blockData", "test")

	request := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(data.Encode())) // URL-encoded payload
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	responseRecorder := httptest.NewRecorder()
	des.add(responseRecorder, request)
	assert.Equal(t, responseRecorder.Code, 308)

	request = httptest.NewRequest(http.MethodGet, "/", nil)
	responseRecorder = httptest.NewRecorder()
	des.home(responseRecorder, request)
	assert.Equal(t, responseRecorder.Code, 200)

	// assert.Equal(t, responseRecorder.Body.String(), "test")
}

func TestClose(t *testing.T) {
	assert.Implements(t, (*explorerserver.ExplorerServer)(nil), new(DefaultExplorerServer),
		"It must assert.Implements of interface explorerserver.ExplorerServer")
}

func initTest() (viperjacket.Config, blockchain.Blockchain, explorerserver.ExplorerServer, error) {
	cfg := wirejacket.GetConfig()

	bc := &mocks.BlockchainMock{}
	bc.Init()

	es := New(cfg, bc)

	return cfg, bc, es, nil
}

func closeTest(cfg viperjacket.Config, bc blockchain.Blockchain, es explorerserver.ExplorerServer) error {
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
