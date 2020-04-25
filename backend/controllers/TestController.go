package controllers

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camStructs"
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
	cam.App.Trace("title", "content")
	cam.App.Debug("title", "content")
	cam.App.Info("title", "content")
	cam.App.Warn("title", "content")
	cam.App.Error("title", "content")
	cam.App.Fatal("title", "content")
	ctrl.SetResponse([]byte("done"))
}

// cache test
func (ctrl *TestController) Cache() {
	cache := cam.App.GetCache()

	msg := new(camStructs.Message)
	msg.Data = "123123123"
	msg.Id = 123
	msg.Route = "abc/xyz"
	cache.SetDuration("short", msg, 100*time.Minute)
	cache.Set("tt", "123123")
	v := cache.Get("tt")
	str, ok := v.(string)
	if !ok {
		str = "fail"
	}
	cam.Debug("TestController.Cache", "v = "+str)

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

func (ctrl *TestController) Post() {
	abc, _ := ctrl.GetValue("abc").(string)
	cam.App.Info("TestController.Post", "acb = "+abc)
}

func (ctrl *TestController) Socket() {
	ctrl.SetResponse([]byte("123123"))
}

func (ctrl *TestController) Valid() {
	v := new(structs.Valid)
	v.Email = "123@123.com"
	v.MyEmail = "abc@abc.a"

	err, _ := cam.Valid(v)
	if err != nil {
		cam.Debug("TestController.valid", err.Error())
	} else {
		cam.Debug("TestController.valid", "no error")
	}
}
