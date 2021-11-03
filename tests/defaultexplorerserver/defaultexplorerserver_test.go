package defaultexplorerserver

import (
	"testing"

	"github.com/bang9211/ossicones/implements/defaultexplorerserver"
	"github.com/bang9211/ossicones/interfaces/explorerserver"

	. "github.com/stretchr/testify/assert"
)

func TestImplementExplorerServer(t *testing.T) {
	Implements(t, (*explorerserver.ExplorerServer)(nil), new(defaultexplorerserver.DefaultExplorerServer),
		"It must implements of interface explorerserver.ExplorerServer")
}

func TestServe(t *testing.T) {
	Implements(t, (*explorerserver.ExplorerServer)(nil), new(defaultexplorerserver.DefaultExplorerServer),
		"It must implements of interface explorerserver.ExplorerServer")
}

func TestClose(t *testing.T) {
	Implements(t, (*explorerserver.ExplorerServer)(nil), new(defaultexplorerserver.DefaultExplorerServer),
		"It must implements of interface explorerserver.ExplorerServer")
}
