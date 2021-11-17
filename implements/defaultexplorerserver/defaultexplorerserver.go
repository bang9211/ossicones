package defaultexplorerserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/explorerserver"
	viperjacket "github.com/bang9211/viper-jacket"
)

const (
	hostConfigPath         = "ossicones_explorer_server_host"
	portConfigPath         = "ossicones_explorer_server_port"
	templatePathConfigPath = "ossicones_explorer_server_template_path"
	defaultHost            = "0.0.0.0"
	defaultPort            = 3000
	defaultTemplatePath    = "${OSSICONES_SRC_HOME}/templates"

	templatePagePath    = "/pages/*.gohtml"
	templatePartialPath = "/partials/*.gohtml"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []blockchain.Block
}

type DefaultExplorerServer struct {
	config       viperjacket.Config
	handler      *http.ServeMux
	blockchain   blockchain.Blockchain
	templatePath string
	address      string
}

// New creates, initializes and returns DefaultExplorerServer.
func New(
	config viperjacket.Config,
	blocchain blockchain.Blockchain) explorerserver.ExplorerServer {
	des := &DefaultExplorerServer{
		config:     config,
		handler:    http.NewServeMux(),
		blockchain: blocchain,
	}

	err := des.init()
	if err != nil {
		log.Printf("failed to init DefaultExplorerServer : %s", err)
		return nil
	}

	return des
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

	go d.Serve()

	return nil
}

func (d *DefaultExplorerServer) home(rw http.ResponseWriter, r *http.Request) {
	blocks := []blockchain.Block{}
	tempBlocks, err := d.blockchain.AllBlocks()
	if err != nil {
		log.Printf("failed to get all blocks : %s", err)
	} else {
		blocks = append(blocks, tempBlocks...)
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
		err := d.blockchain.AddBlock(data)
		if err != nil {
			fmt.Printf("failed to add block : %s", err)
		}
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func (d *DefaultExplorerServer) Serve() error {
	log.Printf("Listening Explorer Server on %s\n", d.address)
	err := http.ListenAndServe(d.address, d.handler)
	if err != nil {
		log.Printf("failed to ListenAndServe DefaultExplorerServer : %s", err)
		return err
	}
	return nil
}

func (d *DefaultExplorerServer) Close() error {
	return nil
}
