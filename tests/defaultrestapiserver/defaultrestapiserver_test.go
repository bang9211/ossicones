package defaultrestapiserver

import (
	"testing"

	"github.com/bang9211/ossicones/implements/defaultrestapiserver"
	"github.com/bang9211/ossicones/interfaces/restapiserver"

	. "github.com/stretchr/testify/assert"
)

func TestImplementRESTAPIServer(t *testing.T) {
	Implements(t, (*restapiserver.RESTAPIServer)(nil), new(defaultrestapiserver.DefaultRESTAPIServer),
		"It must implements of interface restapiserver.RESTAPIServer")
}

func TestServe(t *testing.T) {
	Implements(t, (*restapiserver.RESTAPIServer)(nil), new(defaultrestapiserver.DefaultRESTAPIServer),
		"It must implements of interface restapiserver.RESTAPIServer")
}

func TestClose(t *testing.T) {
	Implements(t, (*restapiserver.RESTAPIServer)(nil), new(defaultrestapiserver.DefaultRESTAPIServer),
		"It must implements of interface restapiserver.RESTAPIServer")
}
