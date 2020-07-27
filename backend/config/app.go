package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/component/camGRpcClient"
	"github.com/go-cam/cam/component/camGRpcServer"
	"github.com/go-cam/cam/component/camGrpc"
	"google.golang.org/grpc"
	backend_grpc "test/backend-grpc"
	"test/backend/controllers"
	"test/backend/services"
	"test/backend/structs"
	"time"
)

// 获取默认配置
func GetApp() camBase.AppConfigInterface {
	config := cam.NewConfig()
	config.ComponentDict = map[string]camBase.ComponentConfigInterface{
		"ws":               websocketServer(),
		"http":             httpServer(),
		"db":               cam.NewDatabaseConfig("mysql", "127.0.0.1", "3306", "cam", "root", "123456"),
		"console":          cam.NewConsoleConfig(),
		"log":              log(),
		"cache":            cacheConfig(),
		"mail":             mailConfig(),
		"tcp":              socketConfig(),
		"grpc-client":		gRpcClientConfig(),
		"grpc-server":		gRpcServerConfig(),
	}
	return config
}

func log() camBase.ComponentConfigInterface {
	config := cam.NewLogConfig()
	config.PrintLevel = cam.LogLevelAll
	config.WriteLevel = cam.LogLevelAll
	return config
}

func websocketServer() camBase.ComponentConfigInterface {
	config := cam.NewWebsocketConfig(20012)
	config.Register(&controllers.TestController{})
	//config.AddMiddleware("", &middlewares.LogMiddleware{})
	//config.AddMiddleware("", &middlewares.AMiddleware{})
	//config.AddMiddleware("", &middlewares.BMiddleware{})
	return config
}

func httpServer() camBase.ComponentConfigInterface {
	config := cam.NewHttpConfig(20000)
	config.SetContextStruct(&structs.HttpContextAo{})

	config.Register(&controllers.TestController{})
	config.Register(&controllers.FileController{})
	//config.AddMiddleware("", &middlewares.LogMiddleware{})
	//config.AddMiddleware("", &middlewares.AMiddleware{})
	//config.AddMiddleware("", &middlewares.BMiddleware{})
	return config
}

func cacheConfig() camBase.ComponentConfigInterface {
	config := cam.NewCacheConfig()
	cache := cam.NewRedisCache()
	cache.SetBase64Crypt()
	config.Engine = cache
	return config
}

func mailConfig() camBase.ComponentConfigInterface {
	account := cam.App.GetEvn("EMAIL_ACCOUNT")
	password := cam.App.GetEvn("EMAIL_PASSWORD")
	host := cam.App.GetEvn("EMAIL_HOST")
	config := cam.NewMailConfig(account, password, host)
	return config
}

func socketConfig() camBase.ComponentConfigInterface {
	config := cam.NewSocketConfig(20022)
	config.Trace = true

	config.Register(&controllers.TestController{})
	config.Register(&controllers.FileController{})
	//config.AddMiddleware("", &middlewares.LogMiddleware{})
	//config.AddMiddleware("", &middlewares.AMiddleware{})
	//config.AddMiddleware("", &middlewares.BMiddleware{})
	return config
}

func grpcBackendCliConfig() camBase.ComponentConfigInterface {
	conf := camGrpc.NewGRpcConfig(&camGrpc.Option{
		Type: camGrpc.TypeClient,
		Client: camGrpc.ClientOption{
			ServerList: []string{"localhost:30001"},
		},
	})
	return conf
}

func grpcBackendSrvConfig() camBase.ComponentConfigInterface {
	conf := camGrpc.NewGRpcConfig(&camGrpc.Option{
		Type: camGrpc.TypeServer,
		Server: camGrpc.ServerOption{
			Addr: "localhost:30001",
			Opts: []grpc.ServerOption{
				grpc.ConnectionTimeout(5 * time.Second),
			},
		},
	})
	conf.RegisterServer(&services.HelloWorldService{})
	return conf
}

func gRpcClientConfig() camBase.ComponentConfigInterface {
	conf := camGRpcClient.NewGRpcClient()
	conf.SetOption(&camGRpcClient.Option{
		Servers: []*camGRpcClient.Server{
			{Addr: "127.0.0.1:30001"},
		},
	})
	return conf
}

func gRpcServerConfig() camBase.ComponentConfigInterface {
	conf := camGRpcServer.NewGRpcServer()
	conf.Port = 30001
	conf.Register(&services.HelloWorldService{}, backend_grpc.RegisterHelloWorldServer)
	return conf
}
