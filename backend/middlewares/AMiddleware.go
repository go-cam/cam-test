package middlewares

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camStatics"
)

type AMiddleware struct {
}

func (mid *AMiddleware) Handler(ctx camStatics.ContextInterface, next camStatics.NextHandler) []byte {
	cam.Debug("AMiddleware", "before")
	res := next()
	cam.Debug("AMiddleware", "after")
	return res
}
