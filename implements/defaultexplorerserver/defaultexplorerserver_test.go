package defaultexplorerserver

import (
	"net/http"
	"net/http/httptest"
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
	t.Setenv("OSSICONES_BLOCKCHAIN_GENESIS_BLOCK_DATA", genesisBlockData)
	cfg, bc, es, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer closeTest(cfg, bc, es)

	des, ok := es.(*defaultexplorerserver.DefaultExplorerServer)
	True(t, ok)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	responseRecorder := httptest.NewRecorder()
	des.home(responseRecorder, request)
}

func TestClose(t *testing.T) {
	Implements(t, (*explorerserver.ExplorerServer)(nil), new(defaultexplorerserver.DefaultExplorerServer),
		"It must implements of interface explorerserver.ExplorerServer")
}
