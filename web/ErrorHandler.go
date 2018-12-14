package web

import (
	"github.com/kataras/iris"
)

func InitErrorHandle() {

	server.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Write("CUSTOM 500 INTERNAL SERVER ERROR PAGE")
		// or ctx.Render, ctx.HTML any render method you want
		ctx.Log("http status: 500 happened!")
	})

	server.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Write("CUSTOM 404 NOT FOUND ERROR PAGE")
		ctx.Log("http status: 404 happened!")
	})
}
