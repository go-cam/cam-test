package controllers

import (
	"github.com/go-cam/cam"
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

// private func
func (controller *TestController) privateFunc() {

}
