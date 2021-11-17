package defaultrestapiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/restapiserver"
	"github.com/bang9211/ossicones/mocks"
	"github.com/bang9211/ossicones/utils"
	viperjacket "github.com/bang9211/viper-jacket"
	wirejacket "github.com/bang9211/wire-jacket"
	"github.com/stretchr/testify/assert"
)

const genesisBlockData = "TEST_GENESIS_BLOCK_DATA"

func TestImplementRESTAPIServer(t *testing.T) {
	assert.Implements(t, (*restapiserver.RESTAPIServer)(nil), new(DefaultRESTAPIServer),
		"It must implements of interface restapiserver.RESTAPIServer")
}

func TestGetRoot(t *testing.T) {
	_, err := utils.GetOrSetHomePath()
	assert.NoError(t, err)
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, rs, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, bc, rs)

	drs, ok := rs.(*DefaultRESTAPIServer)
	assert.True(t, ok)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	responseRecorder := httptest.NewRecorder()
	drs.documentation(responseRecorder, request)
	assert.Equal(t, responseRecorder.Code, 200)
}

func TestPostBlocks(t *testing.T) {
	_, err := utils.GetOrSetHomePath()
	assert.NoError(t, err)
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, rs, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, bc, rs)

	drs, ok := rs.(*DefaultRESTAPIServer)
	assert.True(t, ok)

	jsonStr := []byte(`{"message":"TEST_BLOCK_DATA"}`)
	request := httptest.NewRequest(http.MethodPost, "/blocks", bytes.NewBuffer(jsonStr))
	responseRecorder := httptest.NewRecorder()
	drs.blocks(responseRecorder, request)
	assert.Equal(t, responseRecorder.Code, 201)
}

func TestGetBlocks(t *testing.T) {
	_, err := utils.GetOrSetHomePath()
	assert.NoError(t, err)
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, rs, err := initTest()
	assert.NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, bc, rs)

	drs, ok := rs.(*DefaultRESTAPIServer)
	assert.True(t, ok)

	err = bc.AddBlock("test1")
	assert.NoError(t, err)
	err = bc.AddBlock("test2")
	assert.NoError(t, err)

	request := httptest.NewRequest(http.MethodGet, "/blocks", nil)
	responseRecorder := httptest.NewRecorder()
	drs.blocks(responseRecorder, request)

	assert.Equal(t, responseRecorder.Code, 200)
	var blocks []ossiconesblockchain.OssiconesBlock
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &blocks)
	assert.NoError(t, err)

	assert.Equal(t, 3, len(blocks))
	assert.Equal(t, genesisBlockData, blocks[0].GetData())
	assert.Equal(t, "test1", blocks[1].GetData())
	assert.Equal(t, "test2", blocks[2].GetData())
}

func initTest() (viperjacket.Config, blockchain.Blockchain, restapiserver.RESTAPIServer, error) {
	cfg := wirejacket.GetConfig()

	bc := &mocks.BlockchainMock{}
	bc.Init()

	rs := New(cfg, bc)

	return cfg, bc, rs, nil
}

func closeTest(cfg viperjacket.Config, bc blockchain.Blockchain, rs restapiserver.RESTAPIServer) error {
	err := rs.Close()
	if err != nil {
		return err
	}
	err = bc.Close()
	if err != nil {
		return err
	}

	return cfg.Close()
}
