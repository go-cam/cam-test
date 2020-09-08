package controllers

import (
	"context"
	"fmt"
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camStructs"
	"github.com/go-cam/cam/base/camUtils"
	backend_grpc "test/backend-grpc"
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
	ctrl.GetContext().Write([]byte("done"))
}

// cache test
func (ctrl *TestController) Cache() {
	cache := cam.App.GetCache()

	msg := new(camStructs.Message)
	msg.Data = map[string]interface{}{"testValue": "123123123"}
	msg.Id = 123
	msg.Route = "abc/xyz"
	_ = cache.SetDuration("short", msg, 100*time.Minute)
	_ = cache.Set("tt", "123123")
	v := cache.Get("tt")
	str, ok := v.(string)
	if !ok {
		str = "fail"
	}
	cam.Debug("TestController.Cache", "v = "+str)

	ctrl.GetContext().Write([]byte("cache test done"))
}
func (ctrl *TestController) CacheGet() {
	v := cam.App.GetCache().Get("short")
	if v != nil {
		cam.App.Debug("TestController.CacheGet.has", camUtils.Json.EncodeStr(v))
	} else {
		cam.App.Debug("TestController.CacheGet not", "")
	}

	ctrl.GetContext().Write([]byte("CacheGet test done"))
}

func (ctrl *TestController) CacheFlush() {
	_ = cam.App.GetCache().Flush()
	ctrl.GetContext().Write([]byte("CacheFlush test done"))
}

func (ctrl *TestController) SendMail() {
	mailCompI := cam.App.GetMail()
	err := mailCompI.Send("test", "cam test", "281093509@qq.com")
	if err != nil {
		cam.App.Error("TestController.SendMail", err.Error())
	}
	ctrl.GetContext().Write([]byte("SendMail done,"))
}

func (ctrl *TestController) Recover() {
	rec := ctrl.GetRecover()
	cam.App.Warn("TestController.Recover", rec.Error())
}

func (ctrl *TestController) Panic() {
	panic(camStructs.NewRecover("test"))
}

func (ctrl *TestController) Post() {
	abc, _ := ctrl.GetValue("abc").(string)
	cam.App.Info("TestController.Post", "acb = "+abc)
}

func (ctrl *TestController) Socket() {
	ctrl.GetContext().Write([]byte("123123"))
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

func (ctrl *TestController) Grpc() {
	conn := cam.App.GetGrpcClientConn("grpc-client")
	helloSrv := backend_grpc.NewHelloWorldClient(conn)
	recv, err :=helloSrv.SayHello(context.Background(), &backend_grpc.HelloWorld_SayHello_Recv{
		Name: "123123",
	})
	if err != nil {
		cam.Error("TestController.Grpc", err.Error())
	}
	name := "nn"
	if recv != nil {
		name = recv.GetName()
	}
	fmt.Println(recv)

	ctrl.GetContext().Write([]byte(name))
}
