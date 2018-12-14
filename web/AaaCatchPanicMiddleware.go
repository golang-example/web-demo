package web

import (
	"github.com/kataras/iris"
	"web-demo/util"
)

func init() {
	RegisterPreMiddleware(catchPanicMiddleware{})
}

type catchPanicMiddleware struct {
	// your 'stateless' fields here
}

func (m catchPanicMiddleware) Serve(ctx *iris.Context) {
	defer util.HandelPanic(ctx)

	ctx.Next()
}
