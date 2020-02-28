package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/camBase"
	"test/backend/controllers"
)

// 获取默认配置
func GetApp() *cam.Config {
	config := cam.NewConfig()
	config.ComponentDict = map[string]camBase.ConfigComponentInterface{
		"ws":      websocketServer(),
		"http":    httpServer(),
		"db":      cam.NewDatabaseConfig("mysql", "127.0.0.1", "3306", "cam", "root", "123456"),
		"console": cam.NewConsoleConfig(),
	}
	return config
}

func websocketServer() camBase.ConfigComponentInterface {
	sslCert := cam.App.GetEvn("SSL_CERT")
	sslKey := cam.App.GetEvn("SSL_KEY")
	return cam.NewWebsocketServerConfig(20010).ListenSsl(20011, sslCert, sslKey)
}

func httpServer() camBase.ConfigComponentInterface {
	config := cam.NewHttpServerConfig(20000)
	config.SessionName = "test"
	sslCert := cam.App.GetEvn("SSL_CERT")
	sslKey := cam.App.GetEvn("SSL_KEY")
	config.ListenSsl(20001, sslCert, sslKey)
	config.Register(&controllers.TestController{})
	return config
}
