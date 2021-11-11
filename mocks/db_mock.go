package mocks

import (
	"github.com/stretchr/testify/mock"
)

type DBMock struct {
	mock.Mock
	Blocks map[string][]byte
}

func (d *DBMock) SaveBlock(hash string, data []byte) error {
	d.Blocks[hash] = data
	return nil
}

func (d *DBMock) SaveBlockchain(data []byte) error { return nil }
func (d *DBMock) GetBlockchain() ([]byte, error)   { return nil, nil }

func (d *DBMock) GetBlock(hash string) ([]byte, error) {
	return d.Blocks[hash], nil
}

func (d *DBMock) Close() error { return nil }
