package modules

import (
	"fmt"
	"log"

	"github.com/bang9211/ossicones/interfaces/config"
)

type module interface {
	Close() error
}

var modules = map[string]module{}

// todo
func setDependencies(config config.Config) {

}

// Init injects dependencies and inits of all modules.
func InitModules(homePath string) {
	fmt.Println("Init Modules")

	config, err := InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	config.Load()

	setDependencies(config)

	bc, err := InitBlockchain(config)
	if err != nil {
		log.Fatal(err)
	}
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Thrid Block")
	// bc.PrintBlock()

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
