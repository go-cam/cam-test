package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
	"test/console/controllers"
)

func GetApp() *cam.Config {
	config := cam.NewConfig()
	config.ComponentDict = map[string]camBase.ComponentConfigInterface{
		"console": newConsoleConfig(),
	}
	return config
}

func newConsoleConfig() camBase.ComponentConfigInterface {
	config := cam.NewConsoleConfig()
	config.Register(&controllers.ConsoleController{})
	return config
}
