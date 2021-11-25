# gocommon
golang通用的工具包

## logger：使用zap封装，支持分片切割

使用：
```go
import (
    "github.com/modongning/gocommon/logger"
)


func main() {
    logger.InitLogger("../test.log")
    
    logger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
    logger.Errorf("Error fetching URL %s : Error = %s", url, err)
    logger.Debugf("Trying to hit GET request for %s", url)
}
```

## gorm: 封装了创建gorm.DB

```go
import (
    "github.com/modongning/gocommon/logger"
    "gormUtil github.com/modongning/gocommon/gorm"
)

func main() {
    logger.InitLogger("../test.log")
	
    mysqlConnectionPool := gormUtil.GetGormInstance(logger,"root:123456@tcp(127.0.0.1:3306)/db_name?charset=utf8&parseTime=True&loc=Local")
    db := mysqlConnectionPool.Db
	
    db.(*gorm.DB).Where("id=?",1)
}
```