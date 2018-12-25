package html

import (
	"web-demo/web"
	"github.com/kataras/iris"
	. "web-demo/config"
)

func init() {
	web.RegisterHandler("Get", "/webdemo/html/v1/index", index)
	web.RegisterHandler("Get", "/webdemo/html/v1/userinfo", userinfo)
}

func index(ctx *iris.Context) {
	ctx.MustRender("index.html", nil)
	return
}

func userinfo(ctx *iris.Context) {

	ctx.MustRender("user_info.html", map[string]interface{}{
		"Id": 1,
		"UserName": Cfg.Liang,
		"Pwd": "123456",
		})

	return
}
