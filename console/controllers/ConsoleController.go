package controllers

import "github.com/go-cam/cam"

type ConsoleController struct {
	cam.Controller
}

func (controller *ConsoleController) Run() {
	controller.SetResponse([]byte("run success"))
}
