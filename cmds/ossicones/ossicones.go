package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bang9211/ossicones/modules"
	"github.com/bang9211/ossicones/utils"
	wirejacket "github.com/bang9211/wire-jacket"
)

func main() {
	isRunning, err := utils.IsRunning()
	if err != nil {
		log.Fatal("Failed to check the process is already running : %s", err)
	} else if isRunning {
		log.Fatal("The process is already running")
	}

	// Dependency Injector using WireJacket
	//
	// Inject injects default dependency set of Blockchain.
	// It injects dependencies and inits of modules.
	// - config.Config
	// - blockchain.Blockchain
	// - explorerserver.ExplorerServer
	// - restapiserver.RESTAPIServer
	wj := wirejacket.NewWithServiceName("ossicones").
		SetInjectors(modules.Injectors).
		SetEagerInjectors(modules.EagerInjectors)

	if err := wj.DoWire(); err != nil {
		log.Fatal(err)
	}
	defer wj.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
