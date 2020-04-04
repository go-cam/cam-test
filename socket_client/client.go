package main

import (
	"github.com/go-cam/cam/base/camStructs"
	"github.com/go-cam/cam/base/camUtils"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":20022")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	response := new(camStructs.Response)
	response.Code = 1
	response.Message = "test"
	response.Values = map[string]interface{}{
		"a": 1,
		"b": 2,
	}

	message := new(camStructs.Message)
	message.Id = 1
	message.Route = "test/socket"
	message.Data = camUtils.Json.EncodeStr(response)

	send := camUtils.Json.Encode(message)
	log.Println("send = " + string(send))
	send = append(send, '\x17')

	_, err = conn.Write(send)

	var recv []byte
	_, err = conn.Read(recv)
	log.Println("recv = " + string(recv))

	time.Sleep(5 * time.Second)
	log.Println("test finish")
}
