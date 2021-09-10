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
// Otherwise, it creates and returns the object.
func GetOrCreate(
	config config.Config,
	blocchain blockchain.Blockchain) restapiserver.RESTAPIServer {
	if drs == nil {
		once.Do(func() {
			drs = &defaultRESTAPIServer{
				config:     config,
				handler:    mux.NewRouter(),
				blockchain: blocchain,
			}
		})
		err := drs.init()
		if err != nil {
			drs = nil
			return nil
		}
	}

	return drs
}

func (d *defaultRESTAPIServer) init() error {
	var err error
	d.homePath, err = utils.GetOrSetHomePath()
	if err != nil {
		return err
	}
	host := d.config.GetString("ossicones_rest_api_server_host", defaultHost)
	port := d.config.GetInt("ossicones_rest_api_server_port", defaultPort)
	d.address = host + ":" + strconv.Itoa(port)

	d.handler.Use(jsonContentTypeMiddleware)
	d.handler.HandleFunc("/", d.documentation).Methods("GET")
	d.handler.HandleFunc("/blocks", d.blocks).Methods("GET", "POST")
	d.handler.HandleFunc("/blocks/{height:[0-9]+}", d.block).Methods("GET")

	d.Serve()

	return nil
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func (d *defaultRESTAPIServer) documentation(rw http.ResponseWriter, r *http.Request) {
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
			URL:         defaultURL{d.address, "/blocks/{id}"},
			Method:      "GET",
			Description: "See A Block",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func (d *defaultRESTAPIServer) blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(d.blockchain.AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleError(json.NewDecoder(r.Body).Decode(&addBlockBody))
		d.blockchain.AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (d *defaultRESTAPIServer) block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleError(err)
	block, err := d.blockchain.GetBlock(id)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrorNotFound {
		encoder.Encode(ErrorResponse{err.Error()})
	} else {
		encoder.Encode(block)
	}
}

func (d *defaultRESTAPIServer) Serve() {
	go func() {
		fmt.Printf("Listening REST API Server on %s\n", d.address)
		log.Fatal(http.ListenAndServe(d.address, d.handler))
	}()
}

func (d *defaultRESTAPIServer) Close() error {
	drs = nil
	return nil
}
