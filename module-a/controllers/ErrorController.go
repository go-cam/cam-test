package controllers

import "github.com/go-cam/cam/plugin/camRouter"

type ErrorController struct {
	camRouter.Controller
}

func (ctrl *ErrorController) Recover() {
	rec := ctrl.GetRecover()
	ctrl.GetContext().Write([]byte(rec.Error()))
}
