package defaultrestapiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/restapiserver"
	"github.com/bang9211/ossicones/utils"
	"github.com/gorilla/mux"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = 4000
)

var drs *defaultRESTAPIServer
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

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type defaultRESTAPIServer struct {
	config     config.Config
	handler    *mux.Router
	blockchain blockchain.Blockchain
	homePath   string
	address    string
}

// GetOrCreate returns the existing singletone object of DefaultAPIServer.
// Otherwise. it creates and returns the object.
func GetOrCreate(
	config config.Config,
	homePath string,
	blocchain blockchain.Blockchain) restapiserver.RESTAPIServer {
	if drs == nil {
		once.Do(func() {
			drs = &defaultRESTAPIServer{
				config:     config,
				handler:    mux.NewRouter(),
				homePath:   homePath,
				blockchain: blocchain,
			}
			drs.init()
		})
	}

	return drs
}

func (dhs *defaultRESTAPIServer) init() {
	host := dhs.config.GetString("ossicones_rest_api_server_host", defaultHost)
	port := dhs.config.GetInt("ossicones_rest_api_server_port", defaultPort)
	dhs.address = host + ":" + strconv.Itoa(port)

	dhs.handler.Use(jsonContentTypeMiddleware)
	dhs.handler.HandleFunc("/", dhs.documentation).Methods("GET")
	dhs.handler.HandleFunc("/blocks", dhs.blocks).Methods("GET", "POST")
	dhs.handler.HandleFunc("/blocks/{height:[0-9]+}", dhs.block).Methods("GET")
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func (drs *defaultRESTAPIServer) documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         defaultURL{drs.address, "/"},
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         defaultURL{drs.address, "/blocks"},
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         defaultURL{drs.address, "/blocks/{id}"},
			Method:      "GET",
			Description: "See A Block",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func (drs *defaultRESTAPIServer) blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(drs.blockchain.AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleError(json.NewDecoder(r.Body).Decode(&addBlockBody))
		drs.blockchain.AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (drs *defaultRESTAPIServer) block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleError(err)
	block, err := drs.blockchain.GetBlock(id)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrorNotFound {
		encoder.Encode(ErrorResponse{err.Error()})
	} else {
		encoder.Encode(block)
	}
}

func (drs *defaultRESTAPIServer) Serve() {
	go func() {
		fmt.Printf("Listening REST API Server on %s\n", drs.address)
		log.Fatal(http.ListenAndServe(drs.address, drs.handler))
	}()
}

func (drs *defaultRESTAPIServer) Close() {
}
