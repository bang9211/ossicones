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
	"github.com/bang9211/ossicones/modules"
)

const port string = ":4000"

var homePath string

type homeData struct {
	PageTitle string
	Blocks    []blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("%stemplates/home.gohtml", homePath+"/")))

	obc := ossiconesblockchain.ObtainBlockchain()
	tempBlocks := obc.AllBlocks()
	blocks := []blockchain.Block{}
	for _, block := range tempBlocks {
		blocks = append(blocks, block.(blockchain.Block))
	}
	data := homeData{"Home", blocks}
	tmpl.Execute(rw, data)
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

	bc, err := modules.InitBlockchain()
	if err != nil {
		log.Fatal(err)
	}
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Thrid Block")
	bc.PrintBlock()

	http.HandleFunc("/", home)
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
