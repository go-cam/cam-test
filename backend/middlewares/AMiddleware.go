package middlewares

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
)

type AMiddleware struct {
}

func (mid *AMiddleware) Handler(ctx camBase.ContextInterface, next camBase.NextHandler) []byte {
	cam.Debug("AMiddleware", "before")
	res := next()
	cam.Debug("AMiddleware", "after")
	return res
}
