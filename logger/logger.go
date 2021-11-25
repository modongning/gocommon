/*
https://www.topgoer.com/%E9%A1%B9%E7%9B%AE/log/ZapLogger.html

Zap 日志查询初始化
Zap本身不支持切割归档日志文件
要在zap中加入Lumberjack支持： go get -u github.com/natefinch/lumberjack

使用：
logger.InitLogger("../test.log")

logger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
logger.Errorf("Error fetching URL %s : Error = %s", url, err)
logger.Debugf("Trying to hit GET request for %s", url)
*/

package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type Logger struct {
	*zap.SugaredLogger
}

var logger *Logger
var initFlag bool

func GetLogInterface() *Logger{
	return logger
}

/*
InitLogger 初始化日志插件
*/
func InitLogger(logFile string) {
	//io.Writer抽象实现
	writeSyncer := getLogWriter(logFile)
	// 编码器
	encoder := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel), //打印到控制台
		zapcore.NewCore(encoder, zapcore.AddSync(writeSyncer), zapcore.DebugLevel),
	)
	logger = &Logger{
		zap.New(
			core,
			zap.AddCaller(),
		).Sugar(),
	}
	initFlag = true
}

func getLogWriter(logFile string) zapcore.WriteSyncer {
	//加入Lumberjack支持,切割文件
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    1,  //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 5,  //保留旧文件的最大个数
		MaxAge:     30, //保留旧文件的最大天数
		LocalTime:  true,
		Compress:   false, //是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.121"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	logger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	logger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}
