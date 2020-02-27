package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/camBase"
	"test/console/controllers"
)

func GetApp() *cam.Config {
	config := cam.NewConfig()
	config.ComponentDict = map[string]camBase.ConfigComponentInterface{
		"console": newConsoleConfig(),
	}
	return config
}

func newConsoleConfig() camBase.ConfigComponentInterface {
	config := cam.NewConsoleConfig()
	config.Register(&controllers.ConsoleController{})
	return config
}
