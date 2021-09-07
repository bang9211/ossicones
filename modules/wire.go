//go:build !wireinject
// +build !wireinject

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

//
// Dependency Injection List
//
// injection_list stores A(implement) with B(injector) using map.
//
// For wiring name of implment using in config with injector function.
//
// Examples :
//
// - map[string]interface{}{"viperconfig": InjectViperConfig}
//
var injection_list = map[string]interface{}{
	"viperconfig":           InjectViperConfig,
	"ossiconesblockchain":   InjectOssiconesBlockchain,
	"defaultexplorerserver": InjectDefaultExplorerServer,
	"defaultrestapiserver":  InjectDefaultRESTAPIServer,
}

//
// Only dependency wiring be executed in wire.go.
//
// Function Naming Form :
//
// - func Inject{Implement}() ({Interface}, error) {}
//
// Examples :
//
// - func InjectViperConfig() (config.Config, error){}
// - func InjectOssiconesBlockChain(){}
//

// InjectViperConfig injects dependencies and inits of Config.
func InjectViperConfig() (config.Config, error) {
	wire.Build(viperconfig.NewViperConfig)
	return nil, nil
}

// InjectOssiconesBlockchain injects dependencies and inits of Blockchain.
func InjectOssiconesBlockchain(config config.Config) (blockchain.Blockchain, error) {
	wire.Build(ossiconesblockchain.GetOrCreate)
	return nil, nil
}

// InjectDefaultExplorerServer injects dependencies and inits of ExplorerServer.
func InjectDefaultExplorerServer(homePath string, config config.Config, blockchain blockchain.Blockchain) (explorerserver.ExplorerServer, error) {
	wire.Build(defaultexplorerserver.GetOrCreate)
	return nil, nil
}

// InjectDefaultRESTAPIServer injects dependencies and inits of APiServer.
func InjectDefaultRESTAPIServer(homePath string, config config.Config, blockchain blockchain.Blockchain) (restapiserver.RESTAPIServer, error) {
	wire.Build(defaultrestapiserver.GetOrCreate)
	return nil, nil
}
