package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camStatics"
	"test/console/controllers"
)

func GetApp() camStatics.AppConfigInterface {
	config := cam.NewConfig()
	config.ComponentDict = map[string]camStatics.ComponentConfigInterface{
		"console": newConsoleConfig(),
	}
	return config
}

func newConsoleConfig() camStatics.ComponentConfigInterface {
	config := cam.NewConsoleConfig()
	config.Register(&controllers.ConsoleController{})
	return config
}
