package bolt

import (
	"sync"

	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/database"
	"github.com/bang9211/ossicones/utils"
	"github.com/boltdb/bolt"
)

const (
	defaultPath  = "ossicones.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

var bt *bolt.DB
var once sync.Once

func GetOrCreate(config config.Config) database.Database {
	if bt == nil {
		once.Do(func() {
			var err error
			path := config.GetString("ossicones_bolt_path", defaultPath)
			bt, err = bolt.Open(path, 0600, nil)
			utils.HandleError(err)
			err = bt.Update(func(t *bolt.Tx) error {
				_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
				utils.HandleError(err)
				_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
				return err
			})
			utils.HandleError(err)
		})
	}

	return bt
}
