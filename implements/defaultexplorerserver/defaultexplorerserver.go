package defaultexplorerserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"

	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/explorerserver"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = 3000
)

var dhs *defaultExplorerServer
var once sync.Once
var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []blockchain.Block
}

type defaultExplorerServer struct {
	config     config.Config
	handler    *http.ServeMux
	blockchain blockchain.Blockchain
	homePath   string
	address    string
}

// GetOrCreate returns the existing singletone object of DefaultHTTPServer.
// Otherwise. it creates and returns the object.
func GetOrCreate(
	config config.Config,
	homePath string,
	blocchain blockchain.Blockchain) explorerserver.ExplorerServer {
	if dhs == nil {
		once.Do(func() {
			dhs = &defaultExplorerServer{
				config:     config,
				handler:    http.NewServeMux(),
				homePath:   homePath,
				blockchain: blocchain,
			}
			dhs.init()
		})
	}

	return dhs
}

func (dhs *defaultExplorerServer) init() {
	host := dhs.config.GetString("ossicones_explorer_server_host", defaultHost)
	port := dhs.config.GetInt("ossicones_explorer_server_port", defaultPort)
	dhs.address = host + ":" + strconv.Itoa(port)

	templates = template.Must(template.ParseGlob(dhs.homePath + "/templates/pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(dhs.homePath + "/templates/partials/*.gohtml"))

	dhs.handler.HandleFunc("/", dhs.home)
	dhs.handler.HandleFunc("/add", dhs.add)
}

func (dhs *defaultExplorerServer) home(rw http.ResponseWriter, r *http.Request) {
	tempBlocks := ossiconesblockchain.GetOrCreate(dhs.config).AllBlocks()
	blocks := []blockchain.Block{}
	for _, block := range tempBlocks {
		blocks = append(blocks, block.(blockchain.Block))
	}
	data := homeData{"Home", blocks}

	templates.ExecuteTemplate(rw, "home", data)
}

func (dhs *defaultExplorerServer) add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		dhs.blockchain.AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func (dhs *defaultExplorerServer) Serve() {
	go func() {
		fmt.Printf("Listening Explorer Server on %s\n", dhs.address)
		log.Fatal(http.ListenAndServe(dhs.address, dhs.handler))
	}()
}

func (dhs *defaultExplorerServer) Close() {
}
