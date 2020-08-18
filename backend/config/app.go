package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camConfig"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/component/camCache"
	"github.com/go-cam/cam/component/camConsole"
	"github.com/go-cam/cam/component/camDatabase"
	"github.com/go-cam/cam/component/camGRpcClient"
	"github.com/go-cam/cam/component/camGRpcServer"
	"github.com/go-cam/cam/component/camHttp"
	"github.com/go-cam/cam/component/camLog"
	"github.com/go-cam/cam/component/camMail"
	"github.com/go-cam/cam/component/camSocket"
	"github.com/go-cam/cam/component/camWebsocket"
	"google.golang.org/grpc"
	backend_grpc "test/backend-grpc"
	"test/backend/controllers"
	"test/backend/services"
	"test/backend/structs"
)

// 获取默认配置
func GetApp() camStatics.AppConfigInterface {
	config := camConfig.NewConfig()
	config.ComponentDict = map[string]camStatics.ComponentConfigInterface{
		"ws":               websocketServer(),
		"http":             httpServer(),
		"db":               camDatabase.NewDatabaseComponentConfig("mysql", "127.0.0.1", "3306", "cam", "root", "123456"),
		"console":          camConsole.NewConsoleComponentConfig(),
		"log":              log(),
		"cache":            cacheConfig(),
		"mail":             mailConfig(),
		"tcp":              socketConfig(),
		"grpc-client":		gRpcClientConfig(),
		"grpc-server":		gRpcServerConfig(),
	}
	return config
}

func log() camStatics.ComponentConfigInterface {
	config := camLog.NewLogConfig()
	config.PrintLevel = camStatics.LogLevelAll
	config.WriteLevel = camStatics.LogLevelAll
	return config
}

func websocketServer() camStatics.ComponentConfigInterface {
	config := camWebsocket.NewWebsocketComponentConfig(20012)
	config.Register(&controllers.TestController{})
	//config.AddMiddleware("", &middlewares.LogMiddleware{})
	//config.AddMiddleware("", &middlewares.AMiddleware{})
	//config.AddMiddleware("", &middlewares.BMiddleware{})
	return config
}

func httpServer() camStatics.ComponentConfigInterface {
	config := camHttp.NewHttpComponentConfig(20000)
	config.SetContextStruct(&structs.HttpContextAo{})

	config.Register(&controllers.TestController{})
	config.Register(&controllers.FileController{})
	//config.AddMiddleware("", &middlewares.LogMiddleware{})
	//config.AddMiddleware("", &middlewares.AMiddleware{})
	//config.AddMiddleware("", &middlewares.BMiddleware{})
	return config
}

func cacheConfig() camStatics.ComponentConfigInterface {
	config := camCache.NewCacheConfig()
	cache := camCache.NewRedisEngine()
	cache.SetBase64Crypt()
	config.Engine = cache
	return config
}

func mailConfig() camStatics.ComponentConfigInterface {
	account := cam.App.GetEvn("EMAIL_ACCOUNT")
	password := cam.App.GetEvn("EMAIL_PASSWORD")
	host := cam.App.GetEvn("EMAIL_HOST")
	config := camMail.NewMailConfig(account, password, host)
	return config
}

func socketConfig() camStatics.ComponentConfigInterface {
	config := camSocket.NewSocketComponentConfig(20022)
	config.Trace = true

	cam.Info("socketConfig", "sleep log test")

	config.Register(&controllers.TestController{})
	config.Register(&controllers.FileController{})
	//config.AddMiddleware("", &middlewares.LogMiddleware{})
	//config.AddMiddleware("", &middlewares.AMiddleware{})
	//config.AddMiddleware("", &middlewares.BMiddleware{})
	return config
}

func gRpcClientConfig() camStatics.ComponentConfigInterface {
	conf := camGRpcClient.NewGRpcClient()
	conf.SetOption(&camGRpcClient.Option{
		Servers: []*camGRpcClient.Server{
			{Addr: "127.0.0.1:30001", DialOptions: []grpc.DialOption{grpc.WithInsecure()}},
		},
	})
	return conf
}

func gRpcServerConfig() camStatics.ComponentConfigInterface {
	conf := camGRpcServer.NewGRpcServer()
	conf.Port = 30001
	conf.Register(&services.HelloWorldService{}, backend_grpc.RegisterHelloWorldServer)
	return conf
}
