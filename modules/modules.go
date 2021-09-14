package modules

import (
	"fmt"
	"log"

	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/explorerserver"
	"github.com/bang9211/ossicones/interfaces/restapiserver"
	"github.com/bang9211/wire-jacket/wirejacket"
)

// wire-jacket approach
//
// InjectDefaultSet injects default dependency set of Blockchain.
// It injects dependencies and inits of all modules.
// - config.Config
// - blockchain.Blockchain
// - explorerserver.ExplorerServer
// - restapiserver.RESTAPIServer
func Inject() error {

	wj, err := wirejacket.NewWithInjectors(injectors, eagerInjectors)
	if err != nil {
		return err
	}
	if err := wj.DoWire(); err != nil {
		log.Fatal(err)
	}

	_, ok := wj.GetModuleByType((*blockchain.Blockchain)(nil)).(blockchain.Blockchain)
	if !ok {
		return fmt.Errorf("failed to get ossiconesblockchain")
	}

	_, ok = wj.GetModuleByType((*explorerserver.ExplorerServer)(nil)).(explorerserver.ExplorerServer)
	if !ok {
		return fmt.Errorf("failed to get defaultexplorerserver")
	}

	_, ok = wj.GetModuleByType((*restapiserver.RESTAPIServer)(nil)).(restapiserver.RESTAPIServer)
	if !ok {
		return fmt.Errorf("failed to get defaultrestapiserver")
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
	// todo
	fmt.Printf("Closed Modules")
}
