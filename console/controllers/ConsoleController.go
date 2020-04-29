package controllers

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component/camConsole"
)

type ConsoleController struct {
	camConsole.ConsoleController
}

func (ctrl *ConsoleController) Run() {
	ctrl.GetContext().Write([]byte("run success"))
}

func (ctrl *ConsoleController) HasCmd() {
	cmd := ctrl.GetArgv(0)
	if camUtils.Console.HasCommand(cmd) {
		cam.Debug("ConsoleController.HasCmd", "has")
	} else {
		cam.Debug("ConsoleController.HasCmd", "not")
	}

}
