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
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/utils"
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

type urlDescription struct {
	URL         defaultURL `json:"url"`
	Method      string     `json:"method"`
	Description string     `json:"description"`
	Payload     string     `json:"payload,omitempty"`
}

func (u urlDescription) String() string {
	return ""
}

type AddBlockBody struct {
	Message string
}

type defaultAPIServer struct {
	config     config.Config
	serveMux   *http.ServeMux
	blockchain blockchain.Blockchain
	homePath   string
	address    string
}

// GetOrCreate returns the existing singletone object of DefaultAPIServer.
// Otherwise. it creates and returns the object.
func GetOrCreate(
	config config.Config,
	homePath string,
	blocchain blockchain.Blockchain) apiserver.APIServer {
	if das == nil {
		once.Do(func() {
			das = &defaultAPIServer{
				config:     config,
				serveMux:   http.NewServeMux(),
				homePath:   homePath,
				blockchain: blocchain,
			}
			das.init()
		})
	}

	return das
}

func (dhs *defaultAPIServer) init() {
	host := dhs.config.GetString("ossicones_api_server_host", "0.0.0.0")
	port := dhs.config.GetInt("ossicones_api_server_port", 4001)
	dhs.address = host + ":" + strconv.Itoa(port)

	dhs.serveMux.HandleFunc("/", dhs.documentation)
	dhs.serveMux.HandleFunc("/blocks", dhs.blocks)
}

func (das *defaultAPIServer) documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         defaultURL{das.address, "/"},
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         defaultURL{das.address, "/blocks"},
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         defaultURL{das.address, "/blocks/{id}"},
			Method:      "GET",
			Description: "See A Block",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func (das *defaultAPIServer) blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(das.blockchain.AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleError(json.NewDecoder(r.Body).Decode(&addBlockBody))
		das.blockchain.AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func (das *defaultAPIServer) Serve() {
	go func() {
		fmt.Printf("Listening API Server on %s\n", das.address)
		log.Fatal(http.ListenAndServe(das.address, das.serveMux))
	}()
}

func (das *defaultAPIServer) Close() {
}
