package defaulthttpserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"

	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/httpserver"
)

var dhs *defaultHTTPServer
var once sync.Once
var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []blockchain.Block
}

type defaultHTTPServer struct {
	blockchain blockchain.Blockchain
	homePath   string
	address    string
}

// GetOrCreate returns the existing singletone object of DefaultHTTPServer.
// Otherwise. it creates and returns the object.
func GetOrCreate(
	homePath string,
	blocchain blockchain.Blockchain) httpserver.HTTPServer {
	if dhs == nil {
		once.Do(func() {
			dhs = &defaultHTTPServer{
				homePath:   homePath,
				blockchain: blocchain,
			}
			dhs.init()
		})
	}

	return dhs
}

func (dhs *defaultHTTPServer) init() {
	host := "0.0.0.0"
	port := 4000
	dhs.address = host + ":" + strconv.Itoa(port)

	templates = template.Must(template.ParseGlob(dhs.homePath + "/templates/pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(dhs.homePath + "/templates/partials/*.gohtml"))

	http.HandleFunc("/", dhs.home)
	http.HandleFunc("/add", dhs.add)
}

func (dhs *defaultHTTPServer) home(rw http.ResponseWriter, r *http.Request) {
	tempBlocks := ossiconesblockchain.GetOrCreate().AllBlocks()
	blocks := []blockchain.Block{}
	for _, block := range tempBlocks {
		blocks = append(blocks, block.(blockchain.Block))
	}
	data := homeData{"Home", blocks}

	templates.ExecuteTemplate(rw, "home", data)
}

func (dhs *defaultHTTPServer) add(rw http.ResponseWriter, r *http.Request) {
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

func (dhs *defaultHTTPServer) Serve() {
	go func() {
		fmt.Printf("Listening HTTP Server on %s\n", dhs.address)
		log.Fatal(http.ListenAndServe(dhs.address, nil))
	}()
}

func (dhs *defaultHTTPServer) Close() {
}
