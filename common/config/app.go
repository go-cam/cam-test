package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
)

// 获取默认配置
func GetApp() camBase.AppConfigInterface {
	config := cam.NewConfig()
	return config
}
