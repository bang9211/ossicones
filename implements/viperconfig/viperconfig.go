package viperconfig

import (
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/spf13/viper"
)

type ViperConfig struct {
	path string
}

// NewViperConfig returns new ViperConfig.
func NewViperConfig(path string) config.Config {
	return &ViperConfig{path: path}
}

func (vc *ViperConfig) Load() error {
	viper.SetConfigFile(vc.path)

	return nil
}
