package middlewares

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camStatics"
)

type TraceMiddleware struct {
}

func (m *TraceMiddleware) Handler(ctx camStatics.ContextInterface, next camStatics.NextHandler) []byte {
	route := ctx.GetRoute()

	cam.Trace("TraceMiddleware", "Recv "+route)
	res := next()
	cam.Trace("TraceMiddleware", "Send "+route)
	return res
}
