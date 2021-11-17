package defaultrestapiserver

import (
	"testing"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/restapiserver"
	"github.com/bang9211/ossicones/mocks"
	viperjacket "github.com/bang9211/viper-jacket"
	wirejacket "github.com/bang9211/wire-jacket"
	"github.com/stretchr/testify/assert"
)

func TestImplementRESTAPIServer(t *testing.T) {
	assert.Implements(t, (*restapiserver.RESTAPIServer)(nil), new(DefaultRESTAPIServer),
		"It must implements of interface restapiserver.RESTAPIServer")
}

func TestServe(t *testing.T) {
	assert.Implements(t, (*restapiserver.RESTAPIServer)(nil), new(DefaultRESTAPIServer),
		"It must implements of interface restapiserver.RESTAPIServer")
}

func TestClose(t *testing.T) {
	assert.Implements(t, (*restapiserver.RESTAPIServer)(nil), new(DefaultRESTAPIServer),
		"It must implements of interface restapiserver.RESTAPIServer")
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
