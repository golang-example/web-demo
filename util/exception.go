package util

import (
	"web-demo/exception"
	. "web-demo/log"
	"github.com/kataras/iris"
	"runtime/debug"
)

func HandelPanic(ctx *iris.Context) {
	if err := recover(); err != nil {
		headerRid := ctx.RequestHeader("X-Web-Demo-RequestId")
		if headerRid == "" {
			ErrorLog.Error("----", string(debug.Stack()))
		}else {
			ErrorLog.Error(headerRid + "----", string(debug.Stack()))
		}

		ctx.JSON(iris.StatusOK, exception.SystemException())
	}
}
