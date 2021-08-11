package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

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
	bc := ossiconesblockchain.ObtainBlockchain()
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Thrid Block")
	bc.PrintBlock()

	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
