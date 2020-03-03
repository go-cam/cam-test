package controllers

import (
	"github.com/go-cam/cam"
	"time"
)

// text controller
type TestController struct {
	cam.Controller
}

func (controller *TestController) Init() {
	controller.Controller.Init()
}

func (controller *TestController) GetDefaultActionName() string {
	return "Test"
}

// test action
func (controller *TestController) Test() {
	cam.App.Debug("title", "content")
	cam.App.Info("title", "content")
	cam.App.Warn("title", "content")
	cam.App.Error("title", "content")
	controller.SetResponse([]byte("done"))
}

// cache test
func (controller *TestController) Cache() {
	cam.App.GetCache().SetDuration("short", "123123", 100*time.Minute)
	controller.SetResponse([]byte("cache test done"))
}
func (controller *TestController) CacheGet() {
	v := cam.App.GetCache().Get("short")
	if v != nil {
		cam.App.Debug("TestController.CacheGet.has", v.(string))
	} else {
		cam.App.Debug("TestController.CacheGet not", "")
	}

	controller.SetResponse([]byte("CacheGet test done"))
}

func (controller *TestController) CacheFlush() {
	cam.App.GetCache().Flush()
	controller.SetResponse([]byte("CacheFlush test done"))
}
