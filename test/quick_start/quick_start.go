package main

import "github.com/go-cam/cam"

func main() {
	cam.RegisterController(&TestController{})
	cam.RunDefault()
}

type TestController struct {
	cam.Controller
}

func (ctrl *TestController) Test() {
	ctrl.SetResponse([]byte("test done."))
}
