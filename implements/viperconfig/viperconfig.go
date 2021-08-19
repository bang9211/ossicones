package viperconfig

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/utils"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const defaultConfigFile = "ossicones.conf"

type ViperConfig struct {
	viper *viper.Viper
	Path  string
}

// NewViperConfig returns new ViperConfig.
func NewViperConfig() config.Config {
	vc := ViperConfig{viper: viper.New()}
	vc.init()
	return &vc
}

func (vc *ViperConfig) init() {
	vc.setFlags()
	// only use 'config' flag for reading config file path
	vc.Path = vc.GetString("config", defaultConfigFile)
}

func (vc *ViperConfig) setFlags() {
	flag.String("config", defaultConfigFile,
		"Config file(envfile)[default : ossicones.conf]")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	vc.viper.BindPFlags(pflag.CommandLine)
}

// Load loads config file from path, if the same key exists in environment variables
// Viper overwrites value of same key to environment variables. 
// all the keys store to lowercase.
func (vc *ViperConfig) Load() error {
	if !strings.Contains(vc.Path, "/") {
		vc.viper.AddConfigPath(".")
	}
	vc.viper.AddConfigPath(utils.GetFileDir(vc.Path))
	vc.viper.SetConfigName(utils.GetFileNameFromPath(vc.Path))
	vc.viper.SetConfigType("env")
	vc.viper.AutomaticEnv()

	if err := vc.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Failed to find config file default values will be used : %s", err)
		} else {
			log.Fatal(err)
		}
	}

	err := vc.viper.Unmarshal(vc)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (vc *ViperConfig) GetBool(key string, defaultVal bool) bool {
	if vc.viper.IsSet(key) {
		return vc.viper.GetBool(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetString(key string, defaultVal string) string {
	if vc.viper.IsSet(key) {
		return vc.viper.GetString(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetInt(key string, defaultVal int) int {
	if vc.viper.IsSet(key) {
		return vc.viper.GetInt(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetInt32(key string, defaultVal int32) int32 {
	if vc.viper.IsSet(key) {
		return vc.viper.GetInt32(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetInt64(key string, defaultVal int64) int64 {
	if vc.viper.IsSet(key) {
		return vc.viper.GetInt64(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetUint(key string, defaultVal uint) uint {
	if vc.viper.IsSet(key) {
		return vc.viper.GetUint(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetUint32(key string, defaultVal uint32) uint32 {
	if vc.viper.IsSet(key) {
		return vc.viper.GetUint32(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetUint64(key string, defaultVal uint64) uint64 {
	if vc.viper.IsSet(key) {
		return vc.viper.GetUint64(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetFloat64(key string, defaultVal float64) float64 {
	if vc.viper.IsSet(key) {
		return vc.viper.GetFloat64(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetTime(key string, defaultVal time.Time) time.Time {
	if vc.viper.IsSet(key) {
		return vc.viper.GetTime(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetDuration(key string, defaultVal time.Duration) time.Duration {
	if vc.viper.IsSet(key) {
		return vc.viper.GetDuration(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetIntSlice(key string, defaultVal []int) []int {
	if vc.viper.IsSet(key) {
		return vc.viper.GetIntSlice(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetStringSlice(key string, defaultVal []string) []string {
	if vc.viper.IsSet(key) {
		return vc.viper.GetStringSlice(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetStringMap(key string, defaultVal map[string]interface{}) map[string]interface{} {
	if vc.viper.IsSet(key) {
		return vc.viper.GetStringMap(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetStringMapString(key string, defaultVal map[string]string) map[string]string {
	if vc.viper.IsSet(key) {
		return vc.viper.GetStringMapString(key)
	}
	return defaultVal
}

func (vc *ViperConfig) GetStringMapSlice(key string, defaultVal map[string][]string) map[string][]string {
	if vc.viper.IsSet(key) {
		return vc.viper.GetStringMapStringSlice(key)
	}
	return defaultVal
}