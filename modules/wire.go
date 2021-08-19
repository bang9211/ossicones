//go:build wireinject

package modules

import (
	"fmt"
	"log"

	"github.com/bang9211/ossicones/implements/defaultapiserver"
	"github.com/bang9211/ossicones/implements/defaulthttpserver"
	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/implements/viperconfig"
	"github.com/bang9211/ossicones/interfaces/apiserver"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
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

// InitConfig injects dependencies and inits of Config.
func InitConfig() (config.Config, error) {
	wire.Build(viperconfig.NewViperConfig)
	return nil, nil
}

// InitHTTPServer injects dependencies and inits of HTTPServer.
func InitHTTPServer(homePath string, config config.Config, blockchain blockchain.Blockchain) (httpserver.HTTPServer, error) {
	wire.Build(defaulthttpserver.GetOrCreate)
	return nil, nil
}

// InitAPIServer injects dependencies and inits of APiServer.
func InitAPIServer(homePath string, config config.Config, blockchain blockchain.Blockchain) (apiserver.APIServer, error) {
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
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Thrid Block")
	// bc.PrintBlock()

	config, err := InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	config.Load()

	hs, err := InitHTTPServer(homePath, config, bc)
	if err != nil {
		log.Fatal(err)
	}
	hs.Serve()

	as, err := InitAPIServer(homePath, config, bc)
	if err != nil {
		log.Fatal(err)
	}
	as.Serve()

}

// Close closes all modules gracefully.
func Close() {
	//todo
	fmt.Printf("Closed Modules")
}
