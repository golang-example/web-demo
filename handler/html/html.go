package html

import (
	"web-demo/web"
	"github.com/kataras/iris"
)

func init() {
	web.RegisterHandler("Get", "/webdemo/user/v1/index", index)
	web.RegisterHandler("Get", "/webdemo/user/v1/userinfo", userinfo)
}

func index(ctx *iris.Context) {
	ctx.MustRender("index.html", nil)
	return
}

func userinfo(ctx *iris.Context) {

	ctx.MustRender("user_info.html", map[string]interface{}{
		"Id": 1,
		"UserName": "liang",
		"Pwd": "123456",
		})

	return
}
