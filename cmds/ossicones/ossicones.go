package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bang9211/ossicones/modules"
)

func getHomePath() (string, error) {
	homePath := os.Getenv("OSSICONES_SRC_HOME")
	if homePath == "" {
		cmdOut, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
		if err != nil {
			log.Fatal(fmt.Sprintf(`Error on getting the go-kit base path: %s - %s`, err.Error(), string(cmdOut)))
			return "", err
		}
		homePath = strings.TrimSpace(string(cmdOut))
		fmt.Printf("HOME PATH : %s\n", homePath)
		os.Setenv("OSSICONES_SRC_HOME", homePath)
	}
	return homePath, nil
}

func main() {
	homePath, err := getHomePath()
	if err != nil {
		fmt.Printf("Failed to getHomePath")
		log.Fatal(err)
	}
	modules.Init(homePath)
	defer modules.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
