package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/plugin/camRouter"
	"test/module-a/controllers"
	"test/module-a/middlewares"
)

func GetApp() camBase.AppConfigInterface {
	conf := cam.NewConfig()
	conf.ComponentDict = map[string]camBase.ComponentConfigInterface{
		"http":    http(),
		"db":      db(),
		"console": console(),
	}
	return conf
}

func http() camBase.ComponentConfigInterface {
	httpPort := camUtils.C.StringToUint16(cam.Env("HTTP_PORT"))

	conf := cam.NewHttpConfig(httpPort)
	conf.RouterOption(&camRouter.RouterOption{
		DefaultRoute: "hello/world",
		RecoverRoute: "error/recover",
	})

	conf.Register(&controllers.ErrorController{})
	conf.Register(&controllers.HelloController{})

	conf.AddMiddleware("", &middlewares.TraceMiddleware{})
	return conf
}

func db() camBase.ComponentConfigInterface {
	driver := cam.Env("DB_DRIVER")
	host := cam.Env("DB_HOST")
	port := cam.Env("DB_PORT")
	name := cam.Env("DB_NAME")
	username := cam.Env("DB_USERNAME")
	password := cam.Env("DB_PASSWORD")

	return cam.NewDatabaseConfig(driver, host, port, name, username, password)
}

func console() camBase.ComponentConfigInterface {
	return cam.NewConsoleConfig()
}
