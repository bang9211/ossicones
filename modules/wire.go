//go:build wireinject
// +build wireinject

package modules

import (
	"github.com/bang9211/ossicones/implements/bolt"
	"github.com/bang9211/ossicones/implements/defaultexplorerserver"
	"github.com/bang9211/ossicones/implements/defaultrestapiserver"
	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/database"
	"github.com/bang9211/ossicones/interfaces/explorerserver"
	"github.com/bang9211/ossicones/interfaces/restapiserver"
	"github.com/google/wire"
)

//
// Dependency Injection List
//
// injectors stores module_name(key) with injector_func(value) using map.
// For wiring, name of implement using in config with injector function.
//
// Examples :
//
//	var injectors = map[string]interface{}{
// 		"viperconfig": 			InjectViperConfig,
// 		"ossiconesblockchain":	InjectOssiconesBlockchain,
// 	}
//
// 	var eagerInjectors = map[string]interface{}{
// 		"defaultexplorerserver": InjectDefaultExplorerServer,
// 		"defaultrestapiserver":  InjectDefaultRESTAPIServer,
// 	}
//
var Injectors = map[string]interface{}{
	"ossiconesblockchain":   InjectOssiconesBlockchain,
	"defaultexplorerserver": InjectDefaultExplorerServer,
	"defaultrestapiserver":  InjectDefaultRESTAPIServer,
}

var EagerInjectors = map[string]interface{}{}

//
// Dependency wiring should be specify in wire.go.
//
// Inject functions can have several dependency parameters
// and should have two returns(interface, error).
// Only structure type is allowed, non-structure(int, string, ...) is not allowed for injection.
//
// Function Form :
//
// - func Inject{Implement}() {Interface} {}
// - func Inject{Implement}() ({Interface}, error) {}
// - func Inject{Implement}({Interface}) {Interface} {}
// - func Inject{Implement}({Interface}) ({Interface}, error) {}
//
// Examples :
//
// - func InjectOssiconesBlockChain(config config.Config) blockchain.Blockchain {}
// - func InjectOssiconesBlockChain(config config.Config) (blockchain.Blockchain, error) {}
//

// InjectBolt injects dependencies and inits of Database.
func InjectBolt(config config.Config) (database.Database, error) {
	wire.Build(bolt.New)
	return nil, nil
}

// InjectOssiconesBlockchain injects dependencies and inits of Blockchain.
func InjectOssiconesBlockchain(config config.Config, db database.Database) (blockchain.Blockchain, error) {
	wire.Build(ossiconesblockchain.New)
	return nil, nil
}

// InjectDefaultExplorerServer injects dependencies and inits of ExplorerServer.
func InjectDefaultExplorerServer(
	config config.Config,
	blockchain blockchain.Blockchain,
) (explorerserver.ExplorerServer, error) {
	wire.Build(defaultexplorerserver.New)
	return nil, nil
}

// InjectDefaultRESTAPIServer injects dependencies and inits of APiServer.
func InjectDefaultRESTAPIServer(
	config config.Config,
	blockchain blockchain.Blockchain,
) (restapiserver.RESTAPIServer, error) {
	wire.Build(defaultrestapiserver.New)
	return nil, nil
}
