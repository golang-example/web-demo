package web

import (
	. "web-demo/util/threadlocal"
	. "web-demo/log"
	"github.com/kataras/iris"
	. "github.com/jtolds/gls"
	"strconv"
	"time"
)

func init() {
	RegisterPreMiddleware(accessLogMiddleware{})
}

type accessLogMiddleware struct {
	// your 'stateless' fields here
}

func (m accessLogMiddleware) Serve(ctx *iris.Context) {
	//check header
	//request id
	requestId := ctx.RequestHeader("X-Web-Demo-RequestId")
	if requestId == "" {
		requestId = strconv.FormatInt(time.Now().UnixNano(), 10)
	}

	//access log
	AccessLog.Info(requestId + "\t" + ctx.RequestIP() + "\t" + string(ctx.RequestURI()))

	//response requestId
	ctx.Response.Header.Add("X-Web-Demo-RequestId", requestId)

	//do chian
	Mgr.SetValues(Values{Rid: requestId}, func() {
		ctx.Next()
	})

	//rrlog
	RRLog.Info(requestId + "\t" + "-RequestIP:==" + ctx.RequestIP())
	RRLog.Info(requestId + "\t" + "-RequestP:==" + string(ctx.RequestPath(true)))
	RRLog.Info(requestId + "\t" + "-RequestD:==" + string(ctx.Request.Body()))
	RRLog.Info(requestId + "\t" + "-Response:==" + string(ctx.Response.Body()))
}
