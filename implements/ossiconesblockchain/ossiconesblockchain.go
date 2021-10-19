package ossiconesblockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/database"
	"github.com/bang9211/ossicones/utils"
)

const (
	genesisBlockDataConfigPath = "ossicones_blockchain_genesis_block_data"
	defaultGenesisBlockData    = "Genesis OssiconesBlock"
)

var db database.Database

type OssiconesBlockchain struct {
	config     config.Config
	NewestHash string `json:"newstHash"`
	Height     int    `json:"height"`
}

// New creates, initializes and returns OssiconesBlockchain.
func New(config config.Config, database database.Database) blockchain.Blockchain {
	db = database
	obc := &OssiconesBlockchain{config: config}

	err := obc.init()
	if err != nil {
		log.Printf("failed to init : %s", err)
		return nil
	}
	log.Printf("Newest Hash : %s\n Height : %d", obc.NewestHash, obc.Height)

	return obc
}

func (o *OssiconesBlockchain) init() error {
	checkpoint, err := db.GetCheckpoint()
	if err != nil {
		return fmt.Errorf("failed to get checkpoint : %s", err)
	}
	if checkpoint == nil {
		genesisData := o.config.GetString(genesisBlockDataConfigPath, defaultGenesisBlockData)
		err = o.AddBlock(genesisData)
	} else {
		log.Printf("Restoring...")
		err = o.restore(checkpoint)
	}
	return err
}

func (o *OssiconesBlockchain) AddBlock(data string) error {
	newBlock, err := o.createBlock(data, o.NewestHash, o.Height+1)
	if err != nil {
		return err
	}
	o.NewestHash = newBlock.Hash
	o.Height = newBlock.Height
	err = o.persist()
	if err != nil {
		return err
	}

	return nil
}

func (o *OssiconesBlockchain) createBlock(data string, prevHash string, height int) (*OssiconesBlock, error) {
	newBlock := &OssiconesBlock{
		Data:     data,
		PrevHash: prevHash,
		Height:   height,
	}
	payload := newBlock.Data + newBlock.PrevHash + fmt.Sprintf("%d", newBlock.Height)
	newBlock.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	err := newBlock.persist()
	if err != nil {
		return nil, err
	}

	return newBlock, nil
}

func (o *OssiconesBlockchain) persist() error {
	data, err := utils.ToBytes(o)
	if err != nil {
		return err
	}
	db.SaveBlockchain(data)

	return nil
}

func (o *OssiconesBlockchain) restore(data []byte) error {
	err := utils.FromBytes(o, data)
	if err != nil {
		return err
	}
	return nil
}

func (o *OssiconesBlockchain) AllBlocks() ([]interface{}, error) {
	blocks := []interface{}{}
	hashCursor := o.NewestHash
	for {
		block, err := o.GetBlock(hashCursor)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
		if block.GetPrevHash() != "" {
			hashCursor = block.GetPrevHash()
		} else {
			break
		}
	}
	return blocks, nil
}

var ErrNotFound = errors.New("block not found")

func (o *OssiconesBlockchain) GetBlock(hash string) (blockchain.Block, error) {
	blockBytes, err := db.GetBlock(hash)
	if err != nil {
		return nil, err
	}
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &OssiconesBlock{}
	err = block.restore(blockBytes)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (o *OssiconesBlockchain) PrintBlock() {
	blocks, err := o.AllBlocks()
	if err != nil {
		log.Printf("failed to get all blocks : %s", err)
	}
	for i, block := range blocks {
		fmt.Println(i, ":", *block.(*OssiconesBlock))
	}
}

func (o *OssiconesBlockchain) Close() error {
	return nil
}
