package defaultexplorerserver

import (
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/explorerserver"
	"github.com/bang9211/ossicones/modules"
	wirejacket "github.com/bang9211/wire-jacket"
)

func initTest() (config.Config, blockchain.Blockchain, explorerserver.ExplorerServer, error) {
	cfg := wirejacket.GetConfig()

	bc, err := modules.InjectOssiconesBlockchain(cfg)
	if err != nil {
		return nil, nil, nil, err
	}

	err = bc.Reset()
	if err != nil {
		return nil, nil, nil, err
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
