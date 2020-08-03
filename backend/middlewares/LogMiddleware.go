package middlewares

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camStatics"
)

type LogMiddleware struct {
}

func (mid *LogMiddleware) Handler(ctx camStatics.ContextInterface, next camStatics.NextHandler) []byte {
	cam.Warn("LogMiddleware.handler", "test log")
	return next()
}
