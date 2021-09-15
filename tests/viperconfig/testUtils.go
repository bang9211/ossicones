package viperconfig

import (
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/modules"
)

func initTest() (config.Config, error) {
	cfg, err := modules.InjectViperConfig()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func closeTest(cfg config.Config) error {
	return cfg.Close()
}
