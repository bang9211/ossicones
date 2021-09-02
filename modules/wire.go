//go:build wireinject
// +build wireinject

package modules

import (
	"github.com/bang9211/ossicones/implements/defaultexplorerserver"
	"github.com/bang9211/ossicones/implements/defaultrestapiserver"
	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/implements/viperconfig"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/explorerserver"
	"github.com/bang9211/ossicones/interfaces/restapiserver"
	"github.com/google/wire"
)

// TODO : Default set of ossicones
// var DefaultSet = wire.NewSet(
// 	wire.InterfaceValue(new(config.Config), viperconfig.NewViperConfig()),
// 	wire.InterfaceValue(new(blockchain.Blockchain), ossiconesblockchain.GetOrCreate()),
// )

// InitConfig injects dependencies and inits of Config.
func InitConfig() (config.Config, error) {
	wire.Build(viperconfig.NewViperConfig)
	return nil, nil
}

// InitBlockchains injects dependencies and inits of Blockchain.
func InitBlockchain(config config.Config) (blockchain.Blockchain, error) {
	wire.Build(ossiconesblockchain.GetOrCreate)
	return nil, nil
}

// InitHTTPServer injects dependencies and inits of ExplorerServer.
func InitHTTPServer(homePath string, config config.Config, blockchain blockchain.Blockchain) (explorerserver.ExplorerServer, error) {
	wire.Build(defaultexplorerserver.GetOrCreate)
	return nil, nil
}

// InitAPIServer injects dependencies and inits of APiServer.
func InitAPIServer(homePath string, config config.Config, blockchain blockchain.Blockchain) (restapiserver.RESTAPIServer, error) {
	wire.Build(defaultrestapiserver.GetOrCreate)
	return nil, nil
}
