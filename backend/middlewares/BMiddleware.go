package middlewares

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camStatics"
)

type BMiddleware struct {
}

func (mid *BMiddleware) Handler(ctx camStatics.ContextInterface, next camStatics.NextHandler) []byte {
	cam.Debug("BMiddleware", "before")
	res := next()
	cam.Debug("BMiddleware", "after")
	return res
}
