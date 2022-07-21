package log

import (
	"log"
	"os"
)

var (
	LTrace *log.Logger // 记录所有日志
	LInfo  *log.Logger // 重要的信息
	LWarn  *log.Logger // 需要注意的信息
	LError *log.Logger // 非常严重的问题
)

func init() {
	LTrace = log.New(os.Stdout,
		"[TRACE] ",
		log.Ldate|log.Ltime|log.Llongfile)

	LInfo = log.New(os.Stdout,
		"[INFO] ",
		log.Ldate|log.Ltime|log.Llongfile)

	LWarn = log.New(os.Stdout,
		"[WARNING] ",
		log.Ldate|log.Ltime|log.Llongfile)

	LError = log.New(os.Stdout,
		"[ERROR] ",
		log.Ldate|log.Ltime|log.Llongfile)
}

func Trace(i ...interface{}) {
	LTrace.Println(i)
}
func Info(i ...interface{}) {
	LInfo.Println(i)
}
func Warn(i ...interface{}) {
	LWarn.Println(i)
}
func Error(i ...interface{}) {
	LError.Println(i)
}
