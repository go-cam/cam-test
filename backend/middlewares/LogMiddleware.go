package middlewares

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
)

type LogMiddleware struct {
}

func (mid *LogMiddleware) Handler(ctx camBase.ContextInterface, next camBase.NextHandler) []byte {
	cam.Warn("LogMiddleware.handler", "test log")
	return next()
}
