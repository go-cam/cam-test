package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/camBase"
	"github.com/gorilla/websocket"
	"net/http"
	"test/backend/controllers"
)

// 获取默认配置
func GetApp() camBase.AppConfigInterface {
	config := cam.NewConfig()
	config.ComponentDict = map[string]camBase.ComponentConfigInterface{
		"ws":      websocketServer(),
		"http":    httpServer(),
		"db":      cam.NewDatabaseConfig("mysql", "127.0.0.1", "3306", "cam", "root", "123456"),
		"console": cam.NewConsoleConfig(),
		"log":     cam.NewLogConfig(),
	}
	return config
}

func websocketServer() camBase.ComponentConfigInterface {
	config := cam.NewWebsocketConfig(20012)
	config.IsSslOn = true
	config.SslPort = 20013
	config.SslCertFile = cam.App.GetEvn("SSL_CERT")
	config.SslKeyFile = cam.App.GetEvn("SSL_KEY")
	config.Register(&controllers.TestController{})
	config.AddRoute("test/abc", func(conn *websocket.Conn) []byte {
		return []byte("route test succ")
	})
	return config
}

func httpServer() camBase.ComponentConfigInterface {
	config := cam.NewHttpServerConfig(20000)
	config.SessionName = "test"
	config.IsSslOn = true
	config.SslPort = 20001
	config.SslCertFile = cam.App.GetEvn("SSL_CERT")
	config.SslKeyFile = cam.App.GetEvn("SSL_KEY")
	config.Register(&controllers.TestController{})
	config.AddRoute("test/test", func(responseWriter http.ResponseWriter, request *http.Request) {
		_, _ = responseWriter.Write([]byte("route test succ"))
	})
	return config
}
