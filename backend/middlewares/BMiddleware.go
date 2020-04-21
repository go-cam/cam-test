package middlewares

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
)

type BMiddleware struct {
}

func (mid *BMiddleware) Handler(ctx camBase.ContextInterface, next camBase.NextHandler) []byte {
	cam.Debug("BMiddleware", "before")
	res := next()
	cam.Debug("BMiddleware", "after")
	return res
}
