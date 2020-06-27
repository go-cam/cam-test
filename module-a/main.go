package main

import (
	"github.com/go-cam/cam"
	_ "github.com/go-sql-driver/mysql"
	"test/module-a/config"
	_ "test/module-a/database/migrations"
)

func main() {
	cam.App.AddConfig(config.GetApp())
	cam.App.Run()
}
