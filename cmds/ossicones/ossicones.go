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
	//TODO
	// if utils.IsRunning() {
	// 	log.Fatal("The process is already running")
	// }

	homePath, err := utils.GetOrSetHomePath()
	if err != nil {
		fmt.Printf("Failed to obtainHomePath")
		log.Fatal(err)
	}

	// Dependency Injection using Wire
	modules.Init(homePath)
	defer modules.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
