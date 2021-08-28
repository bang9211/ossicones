package ossiconesblockchain

import (
	"fmt"
	"log"
	"testing"

	"github.com/bang9211/ossicones/modules"
	"github.com/bang9211/ossicones/utils"
)

func init() {
	isRunning, err := utils.IsRunning()
	if err != nil {
		fmt.Printf("Failed to check the process is already running")
		log.Fatal(err)
	} else if isRunning {
		log.Fatal("The process is already running")
	}

	homePath, err := utils.GetOrSetHomePath()
	if err != nil {
		fmt.Printf("Failed to GetOrSetHomePath")
		log.Fatal(err)
	}
	// Dependency Injection using Wire
	modules.InitModules(homePath)
}

func close() {
	modules.Close()
}

func TestAddBlock(t *testing.T) {

}

func TestAllBlocks(t *testing.T) {

}

func TestPrintBlock(t *testing.T) {

}

func TestGetBlock(t *testing.T) {

}
