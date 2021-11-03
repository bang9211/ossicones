package bolt

import (
	"testing"

	"github.com/bang9211/ossicones/implements/bolt"
	"github.com/bang9211/ossicones/interfaces/database"

	. "github.com/stretchr/testify/assert"
)

func TestImplementBolt(t *testing.T) {
	Implements(t, (*database.Database)(nil), new(bolt.BoltDB),
		"It must implements of interface database.Database")
}

func TestSaveBlock(t *testing.T) {
	Implements(t, (*database.Database)(nil), new(bolt.BoltDB),
		"It must implements of interface database.Database")
}

func TestSaveBlockchain(t *testing.T) {
	Implements(t, (*database.Database)(nil), new(bolt.BoltDB),
		"It must implements of interface database.Database")
}

func TestGetCheckpoint(t *testing.T) {
	Implements(t, (*database.Database)(nil), new(bolt.BoltDB),
		"It must implements of interface database.Database")
}

func TestGetBlock(t *testing.T) {
	Implements(t, (*database.Database)(nil), new(bolt.BoltDB),
		"It must implements of interface database.Database")
}

func TestClose(t *testing.T) {
	Implements(t, (*database.Database)(nil), new(bolt.BoltDB),
		"It must implements of interface database.Database")
}
