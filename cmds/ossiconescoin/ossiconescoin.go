package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "/templates/"
)

var templates *template.Template

var homePath string

type homeData struct {
	PageTitle string
	Blocks    []blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tempBlocks := ossiconesblockchain.ObtainBlockchain().AllBlocks()
	blocks := []blockchain.Block{}
	for _, block := range tempBlocks {
		blocks = append(blocks, block.(blockchain.Block))
	}
	data := homeData{"Home", blocks}

	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(rw, "add", nil)
}

func main() {
	homePath = os.Getenv("OSSICONES_SRC_HOME")
	if homePath == "" {
		cmdOut, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
		if err != nil {
			log.Fatal(fmt.Sprintf(`Error on getting the go-kit base path: %s - %s`, err.Error(), string(cmdOut)))
			return
		}
		homePath = strings.TrimSpace(string(cmdOut))
		fmt.Printf("HOME PATH : %s\n", homePath)
		os.Setenv("OSSICONES_SRC_HOME", homePath)
	}

	templates = template.Must(template.ParseGlob(homePath + templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(homePath + templateDir + "partials/*.gohtml"))

	bc := ossiconesblockchain.ObtainBlockchain()
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Thrid Block")
	bc.PrintBlock()

	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)

	dir, _ := os.Getwd()
	fmt.Printf("Listening on http://localhost%s in %s\n", port, dir)
	log.Fatal(http.ListenAndServe(port, nil))
}

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func PrintMyPath() {
	fmt.Println(basepath)
}
