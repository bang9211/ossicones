package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bang9211/ossicones/interfaces/explorerserver"
	"github.com/bang9211/ossicones/interfaces/restapiserver"
	"github.com/bang9211/ossicones/modules"
	"github.com/bang9211/ossicones/utils"
	wirejacket "github.com/bang9211/wire-jacket"
)

func main() {
	isRunning, err := utils.IsRunning()
	if err != nil {
		log.Print("Failed to check the process is already running :", err)
	} else if isRunning {
		log.Fatal("The process is already running")
	}

	_, err = utils.GetOrSetHomePath()
	if err != nil {
		log.Print("Failed to GetOrSetHomePath() :", err)
	}

	wj := wirejacket.NewWithServiceName("ossicones").
		SetInjectors(modules.Injectors).
		SetEagerInjectors(modules.EagerInjectors)

	if err := wj.DoWire(); err != nil {
		log.Fatal(err)
	}
	defer wj.Close()

	startMode(wj)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}

func startMode(wj *wirejacket.WireJacket) error {
	config := wirejacket.GetConfig()
	mode := config.GetString("ossicones_mode", "both")
	switch mode {
	case "rest":
		wj.GetModuleByType((*restapiserver.RESTAPIServer)(nil))
	case "html":
		wj.GetModuleByType((*explorerserver.ExplorerServer)(nil))
	case "both":
		wj.GetModuleByType((*restapiserver.RESTAPIServer)(nil))
		wj.GetModuleByType((*explorerserver.ExplorerServer)(nil))
	default:
		return fmt.Errorf("failed to startMode, mode : %s", mode)
	}
	return nil
}
