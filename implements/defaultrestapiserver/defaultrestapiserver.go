package defaultrestapiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/restapiserver"
	"github.com/bang9211/ossicones/utils"
	"github.com/gorilla/mux"
)

const (
	hostConfigPath = "ossicones_rest_api_server_host"
	portConfigPath = "ossicones_rest_api_server_port"
	defaultHost    = "0.0.0.0"
	defaultPort    = 4000
)

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

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type DefaultRESTAPIServer struct {
	config     config.Config
	handler    *mux.Router
	blockchain blockchain.Blockchain
	homePath   string
	address    string
}

// New returns the existing singletone object of DefaultAPIServer.
// Otherwise, it creates and returns the object.
func New(
	config config.Config,
	blocchain blockchain.Blockchain) restapiserver.RESTAPIServer {
	drs := &DefaultRESTAPIServer{
		config:     config,
		handler:    mux.NewRouter(),
		blockchain: blocchain,
	}

	err := drs.init()
	if err != nil {
		log.Printf("failed to init DefaultRESTAPIServer : %s", err)
		return nil
	}

	return drs
}

func (d *DefaultRESTAPIServer) init() error {
	var err error
	d.homePath, err = utils.GetOrSetHomePath()
	if err != nil {
		return err
	}
	host := d.config.GetString(hostConfigPath, defaultHost)
	port := d.config.GetInt(portConfigPath, defaultPort)
	d.address = host + ":" + strconv.Itoa(port)

	d.handler.Use(jsonContentTypeMiddleware)
	d.handler.HandleFunc("/", d.documentation).Methods("GET")
	d.handler.HandleFunc("/blocks", d.blocks).Methods("GET", "POST")
	d.handler.HandleFunc("/blocks/{hash:[a-f0-9]+}", d.block).Methods("GET")

	go d.Serve()

	return nil
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func (d *DefaultRESTAPIServer) documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         defaultURL{d.address, "/"},
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         defaultURL{d.address, "/blocks"},
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         defaultURL{d.address, "/blocks/{hash}"},
			Method:      "GET",
			Description: "See A Block",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func (d *DefaultRESTAPIServer) blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		blocks, err := d.blockchain.AllBlocks()
		if err != nil {
			log.Printf("failed to get all blocks : %s", err)
		}
		json.NewEncoder(rw).Encode(blocks)
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleError(json.NewDecoder(r.Body).Decode(&addBlockBody))
		d.blockchain.AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (d *DefaultRESTAPIServer) block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	block, err := d.blockchain.GetBlock(hash)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrorNotFound {
		encoder.Encode(ErrorResponse{err.Error()})
	} else {
		encoder.Encode(block)
	}
}

func (d *DefaultRESTAPIServer) Serve() error {
	log.Printf("Listening REST API Server on %s\n", d.address)
	err := http.ListenAndServe(d.address, d.handler)
	if err != nil {
		log.Printf("failed to ListenAndServe DefaultExplorerServer : %s", err)
		return err
	}
	return nil
}

func (d *DefaultRESTAPIServer) Close() error {
	return nil
}
