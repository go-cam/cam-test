package structs

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
	"net/http"
)

type Context struct {
	cam.Context
	camBase.HttpContextInterface

	responseWriter http.ResponseWriter
	request        *http.Request
}

func (c *Context) SetHttpResponseWriter(responseWriter http.ResponseWriter) {
	c.responseWriter = responseWriter
}

func (c *Context) GetHttpResponseWriter() http.ResponseWriter {
	return c.responseWriter
}

func (c *Context) SetHttpRequest(request *http.Request) {
	c.request = request
}

func (c *Context) GetHttpRequest() *http.Request {
	return c.request
}
