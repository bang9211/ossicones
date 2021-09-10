package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bang9211/ossicones/modules"
	"github.com/bang9211/ossicones/utils"
)

func main() {
	isRunning, err := utils.IsRunning()
	if err != nil {
		fmt.Printf("Failed to check the process is already running")
		log.Fatal(err)
	} else if isRunning {
		log.Fatal("The process is already running")
	}

	// Dependency Injection using Wire
	// modules.InitModules()
	modules.Inject()
	defer modules.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
