package modules

import (
	"fmt"
	"log"
	"reflect"

	"github.com/bang9211/ossicones/implements/viperconfig"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/bang9211/ossicones/interfaces/config"
	"github.com/bang9211/ossicones/interfaces/explorerserver"
	"github.com/bang9211/ossicones/interfaces/restapiserver"
)

type closable interface {
	Close() error
}

var instance_list = map[string]closable{}

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
func InjectDefaultSet(homePath string) (
	config.Config,
	blockchain.Blockchain,
	explorerserver.ExplorerServer,
	restapiserver.RESTAPIServer,
	error,
) {
	fmt.Println("Init Modules")

	cfg := viperconfig.NewViperConfig()
	activatingModules := readActivatingModules(cfg)

	for _, activatingModule := range activatingModules {
		//get function type
		methodType := reflect.ValueOf(injection_list[activatingModule]).Type()
		fmt.Println(methodType)

		//check dependencies(function params)
		if methodType.NumIn() == 0 {
			//call
			instance := reflect.ValueOf(injection_list[activatingModule]).Call(nil)
			if len(instance) == 2 { // return (instance, error)
				if !instance[1].CanInterface() {
					return nil, nil, nil, nil, fmt.Errorf(
						"[%s] failed to cast error(%s) to interface",
						activatingModule, instance[1],
					)
				}
				err := instance[1].Interface()
				if err != nil {
					return nil, nil, nil, nil, fmt.Errorf(
						"[%s] failed to inject : %s",
						activatingModule, err,
					)
				}

				if !instance[0].CanInterface() {
					return nil, nil, nil, nil, fmt.Errorf(
						"[%s] failed to cast instance(%s) to interface",
						activatingModule, instance[0],
					)
				}
				closableModule, ok := instance[0].Interface().(closable)
				if !ok {
					return nil, nil, nil, nil, fmt.Errorf(
						"[%s] failed to cast instance(%s) to closable",
						activatingModule, instance[0],
					)
				}
				instance_list[activatingModule] = closableModule
			} else if len(instance) == 1 { // return (instance)
				if !instance[0].CanInterface() {
					return nil, nil, nil, nil, fmt.Errorf(
						"[%s] instance(%s) can't be interface",
						activatingModule, instance[0],
					)
				}
				closableModule, ok := instance[0].Interface().(closable)
				if !ok {
					return nil, nil, nil, nil, fmt.Errorf(
						"[%s] failed to cast instance(%s) to closable",
						activatingModule, instance[0],
					)
				}
				instance_list[activatingModule] = closableModule
			} else {
				return nil, nil, nil, nil, fmt.Errorf(
					"[%s] unsupported format of inject function, len(returns) : %d",
					activatingModule, len(instance),
				)
			}
			// instance_list[activatingModule] = .
		} else {
			//store

		}
	}

	bc, err := InjectOssiconesBlockchain(cfg)
	if err != nil {
		log.Fatal(err)
	}
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Thrid Block")
	// bc.PrintBlock()

	hs, err := InjectDefaultExplorerServer(homePath, cfg, bc)
	if err != nil {
		log.Fatal(err)
	}
	hs.Serve()

	as, err := InjectDefaultRESTAPIServer(homePath, cfg, bc)
	if err != nil {
		log.Fatal(err)
	}
	as.Serve()

	return cfg, bc, hs, as, nil
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
