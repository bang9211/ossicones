package wirejacket

import (
	"log"
)

type modulable interface {
	Close() error
}

type WireJacket struct {
	instance_list map[string]modulable
	config        Config
}

func New(wirePkgName string) *WireJacket {
	wj := &WireJacket{
		instance_list: map[string]modulable{},
		config:        NewViperConfig(),
	}

	// activatingModules := readActivatingModules(wj.config)

	return wj
}

func (w *WireJacket) getConfig() Config {
	return w.config
}

func (w *WireJacket) AddProvider(name string, provider interface{}) {

}

func (w *WireJacket) SetProvider(injection_list map[string]interface{}) {

}

func (w *WireJacket) GetInstance(interfaceType *interface{}) interface{} {
	return *interfaceType
}

func readActivatingModules(config Config) []string {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	activatingModules := config.GetStringSlice(
		"ossicones_activating_modules",
		[]string{},
		// defaultActivatingModules[:], // array to slice
	)

	return activatingModules
}
