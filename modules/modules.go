package modules

import (
	"fmt"
	"log"

	wirejacket "github.com/bang9211/wire-jacket"
)

var wj *wirejacket.WireJacket

// wire-jacket approach
//
// InjectDefaultSet injects default dependency set of Blockchain.
// It injects dependencies and inits of all modules.
// - config.Config
// - blockchain.Blockchain
// - explorerserver.ExplorerServer
// - restapiserver.RESTAPIServer
func Inject() error {
	var err error
	wj, err = wirejacket.NewWithInjectors(injectors, eagerInjectors)
	if err != nil {
		return err
	}

	// wj, err := wirejacket.New()
	// if err != nil {
	// 	return err
	// }

	// wj.AddInjector("viperconfig", InjectViperConfig)
	// wj.AddInjector("ossiconesblockchain", InjectOssiconesBlockchain)
	// wj.AddEagerInjector("defaultexplorerserver", InjectDefaultExplorerServer)
	// wj.AddEagerInjector("defaultrestapiserver", InjectDefaultRESTAPIServer)

	if err := wj.DoWire(); err != nil {
		log.Fatal(err)
	}

	return nil
}

// wire approach
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
}

// Close closes all modules gracefully.
func Close() {
	if err := wj.Close(); err != nil {
		log.Fatal(err)
	}
}
