package ossiconesblockchain

import (
	"fmt"
	"os"

	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/database"
	"github.com/bang9211/ossicones/modules"
	wirejacket "github.com/bang9211/wire-jacket"
)

const genesisBlockData = "TEST_GENESIS_BLOCK_DATA"

func initTest() (config.Config, database.Database, blockchain.Blockchain, error) {
	cfg := wirejacket.GetConfig()

	err := os.Remove("ossicones.db")

	db, err := modules.InjectBolt(cfg)
	if err != nil {
		return nil, nil, nil, err
	}

	bc := ossiconesblockchain.New(cfg, db)
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
	return cfg.Close()
}
