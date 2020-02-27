package config

import (
	"github.com/go-cam/cam"
	"test/common/config"
)

// 加载配置
func LoadConfig() {
	// load common's config
	cam.App.AddConfig(config.GetApp())
	// load module's config
	cam.App.AddConfig(GetApp())
}
