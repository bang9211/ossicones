package database

type Database interface {
	// SaveBlock saves the block to DB.
	SaveBlock(hash string, data []byte) error
	// SaveBlockchain saves the blockchain to DB.
	SaveBlockchain(data []byte) error
	// GetBlockchain gets checkpoint from DB.
	GetBlockchain() ([]byte, error)
	// GetBlock gets block from DB.
	GetBlock(hash string) ([]byte, error)
	// Close closes the Explorer Server.
	Close() error
}
