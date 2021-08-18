package defaultapiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/bang9211/ossicones/interfaces/apiserver"
	"github.com/bang9211/ossicones/interfaces/blockchain"
)

var das *defaultAPIServer
var once sync.Once

type defaultURL struct {
	address string
	path    string
}

func (u *defaultURL) String() string {
	return u.path
}

func (u *defaultURL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("%s%s", u.address, u.path)
	return []byte(url), nil
}

type URLDescription struct {
	URL         defaultURL `json:"url"`
	Method      string     `json:"method"`
	Description string     `json:"description"`
	Payload     string     `json:"payload,omitempty"`
}

func (u URLDescription) String() string {
	return ""
}

type defaultAPIServer struct {
	blockchain blockchain.Blockchain
	homePath   string
	address    string
}

// GetOrCreate returns the existing singletone object of DefaultHTTPServer.
// Otherwise. it creates and returns the object.
func GetOrCreate(
	homePath string,
	blocchain blockchain.Blockchain) apiserver.APIServer {
	if das == nil {
		once.Do(func() {
			das = &defaultAPIServer{
				homePath:   homePath,
				blockchain: blocchain,
			}
			das.init()
		})
	}

	return das
}

func (dhs *defaultAPIServer) init() {
	host := "0.0.0.0"
	port := 4001
	dhs.address = host + ":" + strconv.Itoa(port)

	// http.HandleFunc("/", dhs.documentation)
	// http.HandleFunc("/add", dhs.blocks)
}

func (das *defaultAPIServer) documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         defaultURL{das.address, "/"},
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         defaultURL{das.address, "blocks"},
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func (das *defaultAPIServer) blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(das.blockchain.AllBlocks())
	}
}

func (das *defaultAPIServer) Serve() {
	go func() {
		fmt.Printf("Listening API Server on %s\n", das.address)
		log.Fatal(http.ListenAndServe(das.address, nil))
	}()
}

func (das *defaultAPIServer) Close() {
}
