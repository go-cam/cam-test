package controllers

import (
	"github.com/go-cam/cam"
	"test/backend/structs"
	"time"
)

// text controller
type TestController struct {
	cam.Controller
}

func (ctrl *TestController) Init() {
	ctrl.Controller.Init()
}

func (ctrl *TestController) GetDefaultActionName() string {
	return "Test"
}

// test action
func (ctrl *TestController) Test() {
	cam.App.Debug("title", "content")
	cam.App.Info("title", "content")
	cam.App.Warn("title", "content")
	cam.App.Error("title", "content")
	ctrl.SetResponse([]byte("done"))
}

// cache test
func (ctrl *TestController) Cache() {
	cam.App.GetCache().SetDuration("short", "123123", 100*time.Minute)
	ctrl.SetResponse([]byte("cache test done"))
}
func (ctrl *TestController) CacheGet() {
	v := cam.App.GetCache().Get("short")
	if v != nil {
		cam.App.Debug("TestController.CacheGet.has", v.(string))
	} else {
		cam.App.Debug("TestController.CacheGet not", "")
	}

	ctrl.SetResponse([]byte("CacheGet test done"))
}

func (ctrl *TestController) CacheFlush() {
	cam.App.GetCache().Flush()
	ctrl.SetResponse([]byte("CacheFlush test done"))
}

func (ctrl *TestController) SendMail() {
	mailCompI := cam.App.GetMail()
	err := mailCompI.Send("test", "cam test", "281093509@qq.com")
	if err != nil {
		cam.App.Error("TestController.SendMail", err.Error())
	}
	ctrl.SetResponse([]byte("SendMail done,"))
}

func (ctrl *TestController) Recover() {
	rec := ctrl.GetRecover()
	cam.App.Warn("TestController.Recover", rec.Error())
}

func (ctrl *TestController) Panic() {
	panic(cam.NewRecover("test"))
}

func (ctrl *TestController) ContextHttpInterface() {
	ctxI := ctrl.GetContext()
	ctx, ok := ctxI.(*structs.Context)
	if !ok {
		ctrl.SetResponse([]byte("fail"))
		return
	}

	ctrl.SetResponse([]byte(ctx.GetHttpRequest().RemoteAddr))
}
