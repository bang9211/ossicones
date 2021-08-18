//+build wireinject

package modules

import (
	"fmt"
	"log"

	"github.com/bang9211/ossicones/implements/defaultapiserver"
	"github.com/bang9211/ossicones/implements/defaulthttpserver"
	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/apiserver"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/httpserver"
	"github.com/google/wire"
)

// TODO : Default set of ossicones
var DefaultSet = wire.NewSet(
	wire.InterfaceValue(new(blockchain.Blockchain), ossiconesblockchain.GetOrCreate()),
)

// InitBlockchains injects dependencies and inits of Blockchain.
func InitBlockchain() (blockchain.Blockchain, error) {
	wire.Build(ossiconesblockchain.GetOrCreate)
	return nil, nil
}

// InitHTTPServer injects dependencies and inits of HTTPServer.
func InitHTTPServer(homePath string, blockchain blockchain.Blockchain) (httpserver.HTTPServer, error) {
	wire.Build(defaulthttpserver.GetOrCreate)
	return nil, nil
}

// InitAPIServer injects dependencies and inits of APiServer.
func InitAPIServer(homePath string, blockchain blockchain.Blockchain) (apiserver.APIServer, error) {
	wire.Build(defaultapiserver.GetOrCreate)
	return nil, nil
}

// Init injects dependencies and inits of all modules.
func InitModules(homePath string) {
	fmt.Println("Init Modules")

	bc, err := InitBlockchain()
	if err != nil {
		log.Fatal(err)
	}

	hs, err := InitHTTPServer(homePath, bc)
	if err != nil {
		log.Fatal(err)
	}

	as, err := InitAPIServer(homePath, bc)
	if err != nil {
		log.Fatal(err)
	}

	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Thrid Block")
	bc.PrintBlock()

	hs.Serve()
	as.Serve()
}

// Close closes all modules gracefully.
func Close() {
	//todo
	fmt.Printf("Closed Modules")
}
