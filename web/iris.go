package web

import (
	"github.com/kataras/iris"
)

type regInfo struct {
	Method  string
	Uri     string
	Handler []iris.HandlerFunc
}

var _RegInfo []regInfo = make([]regInfo, 0, 50)
var _PreMiddleware iris.Middleware = make([]iris.Handler, 0, 10)

func RegisterHandler(method, uri string, handler ...iris.HandlerFunc) {
	info := regInfo{method, uri, handler}
	_RegInfo = append(_RegInfo, info)
}

func RegisterPreMiddleware(m iris.Handler) {
	_PreMiddleware = append(_PreMiddleware, m)
}

var server *iris.Framework

func RunIris(port string) {
	server = iris.New()

	//init error handle
	InitErrorHandle()

	// add middleware here
	server.Use(_PreMiddleware...)

	ProcessMapping(server, &_RegInfo)

	server.Listen(":" + port)
}

func ProcessMapping(mw iris.MuxAPI, registerInfo *[]regInfo) {
	for _, info := range *registerInfo {
		switch info.Method {
		case "Get":
			mw.Get(info.Uri, info.Handler...)
		case "Post":
			mw.Post(info.Uri, info.Handler...)
		case "Put":
			mw.Put(info.Uri, info.Handler...)
		case "Delete":
			mw.Delete(info.Uri, info.Handler...)
		case "Options":
			mw.Options(info.Uri, info.Handler...)
		}
	}
}
