package bolt

import (
	"os"

	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/database"
	"github.com/bang9211/ossicones/modules"
	wirejacket "github.com/bang9211/wire-jacket"
)

func initTest() (config.Config, database.Database, error) {
	cfg := wirejacket.GetConfig()

	err := os.Remove("ossicones.db")

	db, err := modules.InjectBolt(cfg)
	if err != nil {
		return nil, nil, err
	}

	return cfg, db, nil
}

func closeTest(cfg config.Config, db database.Database) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return cfg.Close()
}
