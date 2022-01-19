package html

import (
	"web-demo/web"
	"github.com/kataras/iris"
	"web-demo/util"
	"fmt"
)

func init() {
	web.RegisterHandler("Get", "/webdemo/html/v1/passenger/login.md", login)
}

func login(ctx *iris.Context) {
	//读取md文件内容
	markdownContents := util.ReadFileToString("templates/passenger/login.md")
	fmt.Println(markdownContents)
	ctx.Markdown(200, markdownContents)
	//页面显示的md内容,缺少了样式 不好看,不行 整合swagger试试
	return
}
