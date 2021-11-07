package bolt

import (
	"log"
	"os"

	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/database"
	"github.com/bang9211/ossicones/utils"
	"github.com/boltdb/bolt"
)

const (
	boltPathConfigPath   = "ossicones_bolt_path"
	defaultBoltPath      = "ossicones.db"
	blockchainBucket     = "blockchain"
	blocksBucket         = "blocks"
	blockchainCheckpoint = "blockchain_checkpoint"
)

type BoltDB struct {
	config config.Config
	bolt   *bolt.DB
}

// New creates, initializes and returns BoltDB.
func New(config config.Config) database.Database {
	bt := &BoltDB{
		config: config,
	}
	err := bt.init()
	if err != nil {
		log.Printf("failed to init BoltDB : %s", err)
		return nil
	}

	return bt
}

func (bdb *BoltDB) init() error {
	boltPath := bdb.config.GetString(boltPathConfigPath, defaultBoltPath)

	dirPath := utils.GetFileDir(boltPath)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	bdb.bolt, err = bolt.Open(boltPath, 0600, nil)
	if err != nil {
		return err
	}
	err = bdb.bolt.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists([]byte(blockchainBucket))
		if err != nil {
			return err
		}
		_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func (bdb *BoltDB) SaveBlockchain(data []byte) error {
	log.Printf("Saving Blockchain")
	// log.Printf("Saving Blockchain : %b", data)
	return bdb.bolt.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blockchainBucket))
		err := bucket.Put([]byte(blockchainCheckpoint), data)
		return err
	})
}

func (bdb *BoltDB) GetBlockchain() ([]byte, error) {
	var data []byte
	err := bdb.bolt.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blockchainBucket))
		data = bucket.Get([]byte(blockchainCheckpoint))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (bdb *BoltDB) SaveBlock(hash string, data []byte) error {
	log.Printf("Saving Block %s", hash)
	// log.Printf("Saving Block %s\nData: %b", hash, data)
	return bdb.bolt.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
}

func (bdb *BoltDB) GetBlock(hash string) ([]byte, error) {
	var data []byte
	err := bdb.bolt.View(func(t *bolt.Tx) error {
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
	return bdb.bolt.Close()
}
