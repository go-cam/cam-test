package controllers

import "github.com/go-cam/cam/plugin/camRouter"

type HelloController struct {
	camRouter.Controller
}

func (ctrl *HelloController) World() {
	name := ctrl.GetString("name")
	if name == "" {
		name = "world"
	}
	ctrl.GetContext().Write([]byte("Hello " + name))
}
