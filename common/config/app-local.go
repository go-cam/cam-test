package config

import (
	"github.com/go-cam/cam"
)

// 获取默认配置
func GetAppLocal() *cam.Config {
	config := cam.NewConfig()
	return config
}
