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
	"github.com/bang9211/ossicones/utils"
	"github.com/bang9211/ossicones/wirejacket"
)

type modulable interface {
	Close() error
}

var instance_list = map[string]modulable{}

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
func Inject() error {
	fmt.Println("Init Modules")

	wj := wirejacket.New("modules")
	wj.SetProvider(injection_list)

	cfg := viperconfig.NewViperConfig()
	activatingModules := readActivatingModules(cfg)
	NotActivatedList := make([]string, len(activatingModules))
	copy(NotActivatedList, activatingModules)

	activatedList := []string{}
	tryCount := 0
	for len(NotActivatedList) > 0 && tryCount < len(NotActivatedList)*len(NotActivatedList) {
		for _, moduleName := range NotActivatedList {
			method := reflect.ValueOf(injection_list[moduleName])
			methodType := method.Type()

			dependencies, satisfied := getNecessaryDependencies(methodType)
			if satisfied {
				returnVal := method.Call(dependencies)
				modulableModule, err := checkInjectionResult(returnVal)
				if err != nil {
					return err
				}
				instance_list[moduleName] = modulableModule
				activatedList = append(activatedList, moduleName)
			}
		}
		for _, activated := range activatedList {
			NotActivatedList = utils.RemoveElement(NotActivatedList, activated)
		}
		tryCount++
	}

	_, ok := instance_list["ossiconesblockchain"].(blockchain.Blockchain)
	if !ok {
		return fmt.Errorf("failed to get ossiconesblockchain")
	}

	_, ok = instance_list["defaultexplorerserver"].(explorerserver.ExplorerServer)
	if !ok {
		return fmt.Errorf("failed to get defaultexplorerserver")
	}

	_, ok = instance_list["defaultrestapiserver"].(restapiserver.RESTAPIServer)
	if !ok {
		return fmt.Errorf("failed to get defaultrestapiserver")
	}

	return nil
}

func getNecessaryDependencies(methodType reflect.Type) ([]reflect.Value, bool) {
	dependencies := []reflect.Value{}
	for i := 0; i < methodType.NumIn(); i++ {
		dependency := methodType.In(i)
		find := false
		for _, instance := range instance_list {
			instanceValue := reflect.ValueOf(instance)
			if instanceValue.CanConvert(dependency) {
				dependencies = append(dependencies, instanceValue)
				find = true
				break
			}
		}
		if !find {
			return nil, false
		}
	}
	return dependencies, true
}

func checkInjectionResult(returnVal []reflect.Value) (modulable, error) {

	if len(returnVal) != 1 && len(returnVal) != 2 {
		return nil, fmt.Errorf(
			"invalid inject function format len(return) : %d", len(returnVal))
	}
	var modulableModule modulable
	var ok bool
	if len(returnVal) == 1 { // return (instance)
		if !returnVal[0].CanInterface() {
			return nil, fmt.Errorf(
				"returnVal(%s) can't be interface",
				returnVal[0],
			)
		}
		modulableModule, ok = returnVal[0].Interface().(modulable)
		if !ok {
			return nil, fmt.Errorf(
				"failed to cast returnVal(%s) to modulable", returnVal[0])
		}
	} else { // return (instance, error)
		if !returnVal[1].CanInterface() {
			return nil, fmt.Errorf(
				"failed to cast error(%s) to interface", returnVal[1])
		}
		err := returnVal[1].Interface()
		if err != nil {
			return nil, fmt.Errorf(
				"failed to inject : %s", err)
		}
		if !returnVal[0].CanInterface() {
			return nil, fmt.Errorf(
				"failed to cast returnVal(%s) to interface", returnVal[0])
		}
		modulableModule, ok = returnVal[0].Interface().(modulable)
		if !ok {
			return nil, fmt.Errorf(
				"failed to cast returnVal(%s) to modulable", returnVal[0])
		}
	}
	return modulableModule, nil
}

func readActivatingModules(config config.Config) []string {

	activatingModules := config.GetStringSlice(
		"ossicones_activating_modules",
		defaultActivatingModules[:], // array to slice
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

	hs, err := InjectDefaultExplorerServer(config, bc)
	if err != nil {
		log.Fatal(err)
	}
	hs.Serve()

	as, err := InjectDefaultRESTAPIServer(config, bc)
	if err != nil {
		log.Fatal(err)
	}
	as.Serve()

}

// Close closes all modules gracefully.
func Close() {
	// todo
	fmt.Printf("Closed Modules")
}
