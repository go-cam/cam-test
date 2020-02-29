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

	controller.SetDefaultActionName("test")
}

// test action
func (controller *TestController) Test() {
	cam.App.Info("title", "content")
	controller.SetResponse([]byte("done"))
}

// private func
func (controller *TestController) privateFunc() {

}
