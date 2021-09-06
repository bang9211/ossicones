package ossiconesblockchain

import (
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/modules"
)

const genesisBlockData = "TEST_GENESIS_BLOCK_DATA"

func initTest() (config.Config, blockchain.Blockchain, error) {
	cfg, err := modules.InitConfig()
	if err != nil {
		return nil, nil, err
	}

	bc, err := modules.InitBlockchain(cfg)
	if err != nil {
		return nil, nil, err
	}

	err = bc.Reset()
	if err != nil {
		return nil, nil, err
	}

	return cfg, bc, nil
}

func closeTest(cfg config.Config, bc blockchain.Blockchain) error {
	err := bc.Close()
	if err != nil {
		return err
	}
	return cfg.Close()
}
