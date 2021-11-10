package ossiconesblockchain

import (
	"fmt"
	"os"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/database"
	wirejacket "github.com/bang9211/wire-jacket"
	"github.com/stretchr/testify/mock"
)

type dbMock struct {
	mock.Mock
	blocks map[string][]byte
}

func (d *dbMock) SaveBlock(hash string, data []byte) error {
	d.blocks[hash] = data
	return nil
}

func (d *dbMock) SaveBlockchain(data []byte) error { return nil }
func (d *dbMock) GetBlockchain() ([]byte, error)   { return nil, nil }

func (d *dbMock) GetBlock(hash string) ([]byte, error) {
	return d.blocks[hash], nil
}

func (d *dbMock) Close() error { return nil }

func initTest() (config.Config, database.Database, blockchain.Blockchain, error) {
	cfg := wirejacket.GetConfig()

	os.Remove("test.db")

	bc := New(cfg, &dbMock{blocks: map[string][]byte{}})
	if bc == nil {
		return nil, nil, nil, fmt.Errorf("failed to New()")
	}

	return cfg, db, bc, nil
}

func closeTest(cfg config.Config, db database.Database, bc blockchain.Blockchain) error {
	err := bc.Close()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}

	os.Remove("ossicones.db")

	return cfg.Close()
}
