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
// Otherwise, it creates and returns the object.
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

func (d *defaultExplorerServer) init() {
	host := d.config.GetString("ossicones_explorer_server_host", defaultHost)
	port := d.config.GetInt("ossicones_explorer_server_port", defaultPort)
	d.address = host + ":" + strconv.Itoa(port)

	templates = template.Must(template.ParseGlob(d.homePath + "/templates/pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(d.homePath + "/templates/partials/*.gohtml"))

	d.handler.HandleFunc("/", d.home)
	d.handler.HandleFunc("/add", d.add)
}

func (d *defaultExplorerServer) home(rw http.ResponseWriter, r *http.Request) {
	tempBlocks := ossiconesblockchain.GetOrCreate(d.config).AllBlocks()
	blocks := []blockchain.Block{}
	for _, block := range tempBlocks {
		blocks = append(blocks, block.(blockchain.Block))
	}
	data := homeData{"Home", blocks}

	templates.ExecuteTemplate(rw, "home", data)
}

func (d *defaultExplorerServer) add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		d.blockchain.AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func (d *defaultExplorerServer) Serve() {
	go func() {
		fmt.Printf("Listening Explorer Server on %s\n", d.address)
		log.Fatal(http.ListenAndServe(d.address, d.handler))
	}()
}

func (d *defaultExplorerServer) Close() error {
	dhs = nil
	return nil
}
