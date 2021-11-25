package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
)

/*
PrometheusBoot
创建普罗米斯监控服务
*/
func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())
	//启动web 服务
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
		if err != nil {
			log.Fatal("启动失败")
		}
		log.Println("监控启动,端口为：" + strconv.Itoa(port))
	}()
}
