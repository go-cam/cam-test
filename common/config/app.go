package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camStatics"
)

// 获取默认配置
func GetApp() camStatics.AppConfigInterface {
	config := cam.NewConfig()
	return config
}
