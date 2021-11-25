package gorm

import (
	logger2 "github.com/modongning/gocommon/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func GetGormInstance(log *logger2.Logger, url string) *MysqlConnectionPool {
	db, err := gorm.Open(
		mysql.Open(url),
		&gorm.Config{
			Logger: logger.New(
				Writer{
					log: log,
				},
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      logger.Info,
					Colorful:      true,
				},
			),
		},
	)
	if err != nil {
		log.Errorf("数据库链接错误：%s", err.Error())

		panic(err.Error())
	}
	return &MysqlConnectionPool{
		db:  db,
		log: log,
	}
}

type MysqlConnectionPool struct {
	db  interface{}
	log *logger2.Logger
}

func (w Writer) Printf(format string, args ...interface{}) {
	w.log.Infof(format, args)
}

type Writer struct {
	log *logger2.Logger
}
