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
	hostConfigPath         = "ossicones_explorer_server_host"
	portConfigPath         = "ossicones_explorer_server_port"
	templatePathConfigPath = "ossicones_explorer_server_template_path"
	defaultHost            = "0.0.0.0"
	defaultPort            = 3000
	defaultTemplatePath    = "templates"

	templatePagePath    = "/pages/*.gohtml"
	templatePartialPath = "/partials/*.gohtml"
)

var dhs *DefaultExplorerServer
var once sync.Once
var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []blockchain.Block
}

type DefaultExplorerServer struct {
	config       config.Config
	handler      *http.ServeMux
	blockchain   blockchain.Blockchain
	templatePath string
	address      string
}

// GetOrCreate returns the existing singletone object of DefaultHTTPServer.
// Otherwise, it creates and returns the object.
func GetOrCreate(
	config config.Config,
	blocchain blockchain.Blockchain) explorerserver.ExplorerServer {
	if dhs == nil {
		once.Do(func() {
			dhs = &DefaultExplorerServer{
				config:     config,
				handler:    http.NewServeMux(),
				blockchain: blocchain,
			}
		})
		err := dhs.init()
		if err != nil {
			dhs = nil
			return nil
		}
	}

	return dhs
}

func (d *DefaultExplorerServer) init() error {
	host := d.config.GetString(hostConfigPath, defaultHost)
	port := d.config.GetInt(portConfigPath, defaultPort)
	d.address = host + ":" + strconv.Itoa(port)
	d.templatePath = d.config.GetString(templatePathConfigPath, defaultTemplatePath)

	templates = template.Must(template.ParseGlob(d.templatePath + templatePagePath))
	templates = template.Must(templates.ParseGlob(d.templatePath + templatePartialPath))

	d.handler.HandleFunc("/", d.home)
	d.handler.HandleFunc("/add", d.add)

	d.Serve()

	return nil
}

func (d *DefaultExplorerServer) home(rw http.ResponseWriter, r *http.Request) {
	tempBlocks := ossiconesblockchain.GetOrCreate(d.config).AllBlocks()
	blocks := []blockchain.Block{}
	for _, block := range tempBlocks {
		blocks = append(blocks, block.(blockchain.Block))
	}
	data := homeData{"Home", blocks}

	templates.ExecuteTemplate(rw, "home", data)
}

func (d *DefaultExplorerServer) add(rw http.ResponseWriter, r *http.Request) {
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

func (d *DefaultExplorerServer) Serve() {
	go func() {
		fmt.Printf("Listening Explorer Server on %s\n", d.address)
		log.Fatal(http.ListenAndServe(d.address, d.handler))
	}()
}

func (d *DefaultExplorerServer) Close() error {
	dhs = nil
	return nil
}
