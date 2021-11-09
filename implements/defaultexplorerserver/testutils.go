package defaultexplorerserver

import (
	"fmt"

	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/explorerserver"
	"github.com/bang9211/ossicones/modules"
	wirejacket "github.com/bang9211/wire-jacket"
)

const genesisBlockData = "TEST_GENESIS_BLOCK_DATA"

func initTest() (config.Config, blockchain.Blockchain, explorerserver.ExplorerServer, error) {
	cfg := wirejacket.GetConfig()

	db, err := modules.InjectBolt(cfg)
	if err != nil {
		return nil, nil, nil, err
	}

	bc := ossiconesblockchain.New(cfg, db)
	if bc == nil {
		return nil, nil, nil, fmt.Errorf("failed to New()")
	}

	es, err := modules.InjectDefaultExplorerServer(cfg, bc)
	if err != nil {
		return nil, nil, nil, err
	}

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
