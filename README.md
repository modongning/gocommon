# gocommon
golang通用的工具包

## logger：使用zap封装，支持分片切割

使用：
```go
import (
    "github.com/modongning/gocommon/logger"
)

var log *logger.Logger

func init() {
	log = logger.NewLogger("./socket.log")
}

func main() {
    log.Infof("启动服务")
}
```

## gorm: 封装了创建gorm.DB

```go
import (
"gormUtil github.com/modongning/gocommon/gorm"
)

func main() {
    mysqlConnectionPool := gormUtil.GetGormInstance(&log,"root:123456@tcp(127.0.0.1:3306)/db_name?charset=utf8&parseTime=True&loc=Local")
    db := mysqlConnectionPool.Db
	
    db.(*gorm.DB).Where("id=?",1)
}
```