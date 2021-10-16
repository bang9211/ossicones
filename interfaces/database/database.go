package database

type Database interface {
	SaveBlock(hash string, data []byte) error
	SaveBlockchain(data []byte) error
	GetCheckpoint() ([]byte, error)
	GetBlock(hash string) ([]byte, error)
}
