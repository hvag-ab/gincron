package logging

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"myproject/pkg/setting"
	"os"
)

//Log配置结构体
type Log struct {
	Gin         string
	App         string
	Http        string
	ServiceName string
}

var LogSetting = &Log{
	Gin:setting.Gin,
	App:setting.Apps,
	Http:setting.Http,
	ServiceName:setting.ServiceName,
}
var AppLogger *zap.Logger//这里定义方便其他包调用 否则被调用出现空指针 首字母也必须大写才能被调用
var HTTPLogger *zap.Logger

var Logger  *zap.Logger
//定制日志
func Setup() {
	// setting.MapTo("log", LogSetting)
	//记录Gin日志
	// f, _ := os.Create(LogSetting.Gin)
	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.DefaultWriter = io.MultiWriter(os.Stdout)//只输入到控制台


	//定制日志 日志等级InfoLevel DebugLeve WarnLevel ErrorLevel DPanicLevel PanicLevel FatalLevel
	AppLogger = NewLogger(LogSetting.App, zapcore.InfoLevel, 128, 30, 7, true, LogSetting.ServiceName)
	HTTPLogger = NewLogger(LogSetting.Http, zapcore.InfoLevel, 128, 30, 7, true, LogSetting.ServiceName)
	/*
	使用方法
	logging.AppLogger.Fatal(fmt.Sprint("Server Shutdown:",err), zap.Error(err))
	logging.AppLogger.Fatal("Server Shutdown:", zap.Error(err)) //logging为包名 *zap.Logger 等级有
	Info Debug Warn Error Dpanic Panic Fatal//如果等级低于设置的等级 那么就不会输出到控制台 也不会保存
	zap.Error(err) 解释 表示格式化输出err错误信息 就不需要自己去字符串格式化了
	zap.String("name",字符串变量) name是键 自己定义  后面是一个string形式的变量
	Zap.Int("key",int型变量)
	zap.Duration("backoff", time.Second)//时间格式化
	*/
	Logger = InitLogger(LogSetting.ServiceName,zap.InfoLevel)//初始化一个日志 只输出到控制台
}
