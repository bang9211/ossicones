//+build wireinject

package modules

import (
	"fmt"
	"log"

	"github.com/bang9211/ossicones/implements/defaulthttpserver"
	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/httpserver"
	"github.com/google/wire"
)

var MySet = wire.NewSet(
	wire.InterfaceValue(
		new(blockchain.Blockchain),
		ossiconesblockchain.ObtainBlockchain(),
	),
)

// var MySet = wire.NewSet(wire.InterfaceValue(new(io.Reader), os.Stdin))

func InitBlockchain() (blockchain.Blockchain, error) {
	wire.Build(MySet)
	return nil, nil
}

func InitHTTPServer(homePath string, blockchain blockchain.Blockchain) (httpserver.HTTPServer, error) {
	wire.Build(defaulthttpserver.ObtainTemplateRouting)
	return nil, nil
}

func Init(homePath string) {
	fmt.Printf("Init Modules")

	bc, err := InitBlockchain()
	if err != nil {
		log.Fatal(err)
	}

	hs, err := InitHTTPServer(homePath, bc)
	if err != nil {
		log.Fatal(err)
	}

	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Thrid Block")
	bc.PrintBlock()

	hs.Serve()
}

func Close() {
	fmt.Printf("Closed Modules")
}
