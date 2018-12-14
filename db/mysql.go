package db

import (
	. "web-demo/log"
	"fmt"
	"github.com/jinzhu/gorm"
	. "web-demo/config"
)

var Mysql *gorm.DB

type SqlDBParam struct {
	Ip       string
	Port     int
	User     string
	Pw       string
	Database string
}
// mysql初始化文件
func SqlDBInit(param *SqlDBParam) {
	Log.Info("init mysql...")
	param_s := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local",
		param.User,
		param.Pw,
		param.Ip,
		param.Port,
		param.Database,
	)
	Log.Info("mysql param: %s", param_s)

	db, err := gorm.Open("mysql", param_s)
	if err != nil {
		Log.Error("open mysql error: %v", err)
		panic(err)
	}
	db.DB().SetMaxIdleConns(Cfg.MysqlMaxConn)
	db.SingularTable(true)
	//db.LogMode(true)

	Mysql = db
	Log.Info("init mysql end.")
}
