/*
https://www.topgoer.com/%E9%A1%B9%E7%9B%AE/log/ZapLogger.html

Zap 日志查询初始化
Zap本身不支持切割归档日志文件
要在zap中加入Lumberjack支持： go get -u github.com/natefinch/lumberjack

使用：
InitLogger("../test.log")

logger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
logger.Errorf("Error fetching URL %s : Error = %s", url, err)
logger.Debugf("Trying to hit GET request for %s", url)
*/

package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args)
}
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args)
}
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args)
}

/*
NewLogger 初始化日志插件
*/
func NewLogger(logFile string) {
	writeSyncer := getLogWriter(logFile)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller()).Sugar()
}

func getLogWriter(logFile string) zapcore.WriteSyncer {
	//加入Lumberjack支持,切割文件
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    1,     //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 5,     //保留旧文件的最大个数
		MaxAge:     30,    //保留旧文件的最大天数
		Compress:   false, //是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
