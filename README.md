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