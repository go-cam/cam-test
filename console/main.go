package main

import (
	"github.com/go-cam/cam"
	"test/console/config"
)

func main() {
	config.LoadConfig()
	cam.App.Run()
}
