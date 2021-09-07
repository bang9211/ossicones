package modules

import (
	"fmt"
	"log"

	"github.com/bang9211/ossicones/implements/viperconfig"
	"github.com/bang9211/ossicones/interfaces/config"
)

var injector wireInjector

type wireInjector struct {
	viperconfig viperconfig.ViperConfig
}
type closable interface {
	Close() error
}

var defaultActivatingModules = [...]string{ //fixed array
	"viperconfig",
	"ossiconesblockchain",
	"defaultexplorerserver",
	"defaultrestapiserver",
}

// InjectDefaultSet injects default dependency set of Blockchain.
// It injects dependencies and inits of all modules.
// - config.Config
// - blockchain.Blockchain
// - explorerserver.ExplorerServer
// - restapiserver.RESTAPIServer
func InjectDefaultSet(homePath string) {
	fmt.Println("Init Modules")

	config := viperconfig.NewViperConfig()
	// activatingModules := readActivatingModules(config)

	// for _, activatingModule := range activatingModules {
	// 	injection_list[activatingModule]
	// }

	bc, err := InjectOssiconesBlockchain(config)
	if err != nil {
		log.Fatal(err)
	}
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Thrid Block")
	// bc.PrintBlock()

	hs, err := InjectDefaultExplorerServer(homePath, config, bc)
	if err != nil {
		log.Fatal(err)
	}
	hs.Serve()

	as, err := InjectDefaultRESTAPIServer(homePath, config, bc)
	if err != nil {
		log.Fatal(err)
	}
	as.Serve()

}

func readActivatingModules(config config.Config) []string {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	activatingModules := config.GetStringSlice(
		"ossicones_activating_modules",
		defaultActivatingModules[:], //array to slice
	)

	return activatingModules
}

func InitModules(homePath string) {
	fmt.Println("Init Modules")

	config, err := InjectViperConfig()
	if err != nil {
		log.Fatal(err)
	}
	config.Load()

	bc, err := InjectOssiconesBlockchain(config)
	if err != nil {
		log.Fatal(err)
	}
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Thrid Block")
	// bc.PrintBlock()

	hs, err := InjectDefaultExplorerServer(homePath, config, bc)
	if err != nil {
		log.Fatal(err)
	}
	hs.Serve()

	as, err := InjectDefaultRESTAPIServer(homePath, config, bc)
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
