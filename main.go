package main

import (
	. "web-demo/config"
	. "web-demo/log"
	. "web-demo/redis"
	. "web-demo/db"
	_ "web-demo/handler"
	"web-demo/web"
	"flag"
	"fmt"
	"time"
)

var showVersion = flag.Bool("v", true, "print version")

func init(){
	//调用 flag.Parse() 进行解析
	flag.Parse()

	//初始化配置文件
	ConfigRead()

	//初始化log日志
	LogInit()

	//初始化redis
	RedisInit(Cfg.RedisAddr,0 , Cfg.RedisPassword, Cfg.RedisMaxConn)

	//初始化mysql
	SqlDBInit(&SqlDBParam{Cfg.MysqlHost, Cfg.MysqlPort, Cfg.MysqlUser, Cfg.MysqlPwd,
	Cfg.MysqlDb})
}

func main() {
	if *showVersion {
		//这个日期就是写死的一个日期，不是这个日期就不认识，就不能正确的格式化
		//据说是go诞生之日
		version := fmt.Sprintf("%s %s@%s", "web-demo", "1.0", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(version)
	}

	Log.Info("start server...")

	//监听端口
	Log.Info("listen on :%s", Cfg.ListenPort)
	web.RunIris(Cfg.ListenPort)
}
