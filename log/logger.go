package log

import (
	"flag"
	"github.com/aiwuTech/fileLogger"
	"os"
	. "web-demo/util/threadlocal"
	"web-demo/config"
)

var Log Logger
var ErrorLog Logger
var AccessLog Logger
var RRLog Logger
var InnerRRLog Logger

type Logger interface {
	Trace(format string, params ...interface{})
	Info(format string, params ...interface{})
	Warn(format string, params ...interface{})
	Error(format string, params ...interface{})
}

type CustomedLogger struct{
	MyLogger Logger
}
func (cl CustomedLogger)Trace(format string, params ...interface{}){
	cl.MyLogger.Trace(cl.resetFormat(format), params)
}
func (cl CustomedLogger)Info(format string, params ...interface{}){
	cl.MyLogger.Info(cl.resetFormat(format), params)
}
func (cl CustomedLogger)Warn(format string, params ...interface{}){
	cl.MyLogger.Warn(cl.resetFormat(format), params)
}
func (cl CustomedLogger)Error(format string, params ...interface{}){
	cl.MyLogger.Error(cl.resetFormat(format), params)
}
func (cl CustomedLogger)resetFormat(format string) string{
	logstr := format
	if rid, ok := Mgr.GetValue(Rid); ok {
		logstr = rid.(string) + " - " + logstr
	}
	return logstr
}

func LogInit() {
	logPath := config.Cfg.LogPath
	if (len(logPath) == 0) {
		logPath = "/Users/liang/ideaWorkspace/go/src/web-demo/logs/web-demo"
	}

	flag.Parse()
	if !isExist(logPath) {
		os.Mkdir(logPath, 0755)
	}

	logger := fileLogger.NewDailyLogger(logPath, "root.log", "", fileLogger.DEFAULT_LOG_SCAN, fileLogger.DEFAULT_LOG_SEQ)
	Log = CustomedLogger{MyLogger: logger}

	errorLog := fileLogger.NewDailyLogger(logPath, "error.log", "", fileLogger.DEFAULT_LOG_SCAN, fileLogger.DEFAULT_LOG_SEQ)
	ErrorLog = CustomedLogger{MyLogger: errorLog}

	accessLog := fileLogger.NewDailyLogger(logPath, "access.log", "", fileLogger.DEFAULT_LOG_SCAN, fileLogger.DEFAULT_LOG_SEQ)
	AccessLog = CustomedLogger{MyLogger: accessLog}

	rRLog := fileLogger.NewDailyLogger(logPath, "rr.log", "", fileLogger.DEFAULT_LOG_SCAN, fileLogger.DEFAULT_LOG_SEQ)
	RRLog = CustomedLogger{MyLogger: rRLog}

	innerRRLog := fileLogger.NewDailyLogger(logPath, "inner_rr.log", "", fileLogger.DEFAULT_LOG_SCAN, fileLogger.DEFAULT_LOG_SEQ)
	InnerRRLog = CustomedLogger{MyLogger: innerRRLog}
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
