package bolt

import (
	"log"
	"sync"

	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/database"
	"github.com/bang9211/ossicones/utils"
	"github.com/boltdb/bolt"
)

const (
	boltPathConfigPath = "ossicones_bolt_path"
	defaultBoltPath    = "ossicones.db"
	dataBucket         = "data"
	blocksBucket       = "blocks"
	checkpoint         = "checkpoint"
)

type BoltDB struct {
	config config.Config
	bolt   *bolt.DB
}

var bt *BoltDB
var once sync.Once

func GetOrCreate(config config.Config) database.Database {
	if bt == nil {
		once.Do(func() {
			bt = &BoltDB{
				config: config,
			}
		})
		err := bt.init()
		if err != nil {
			bt = nil
			return nil
		}
	}

	return bt
}

func (bdb *BoltDB) init() error {
	var err error
	boltPath := bdb.config.GetString(boltPathConfigPath, defaultBoltPath)
	bdb.bolt, err = bolt.Open(boltPath, 0600, nil)
	if err != nil {
		return err
	}
	err = bt.bolt.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
		utils.HandleError(err)
		_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func (bdb *BoltDB) SaveBlock(hash string, data []byte) error {
	log.Printf("Saving Block %s\nData: %b", hash, data)
	return bt.bolt.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
}

func (bdb *BoltDB) SaveBlockchain(data []byte) error {
	log.Printf("Saving Blockchain : %b", data)
	return bt.bolt.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
}

func (bdb *BoltDB) GetCheckpoint() ([]byte, error) {
	var data []byte
	err := bt.bolt.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (bdb *BoltDB) GetBlock(hash string) ([]byte, error) {
	var data []byte
	err := bt.bolt.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (bdb *BoltDB) Close() error {
	return bdb.Close()
}
