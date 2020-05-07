package main

import (
	"fmt"
	_ "git.trac.cn/nv/spider/docs"
	"git.trac.cn/nv/spider/fetcher"
	"git.trac.cn/nv/spider/persist"
	"git.trac.cn/nv/spider/pkg/logging"
	"git.trac.cn/nv/spider/pkg/setting"
	"git.trac.cn/nv/spider/pkg/validator"
	"git.trac.cn/nv/spider/services/spiderhttp/routers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	logging.Setup()
	validator.Setup()
	fetcher.Setup()
	persist.Setup()
}

func main() {
	go func() {
		metricsEndPoint := fmt.Sprintf(":%d", setting.ServerSetting.MetricsPort)
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(metricsEndPoint, nil)
		if err != nil {
			log.Printf("[info] start metrics server of prometheus listening %s", metricsEndPoint)
		}
	}()

	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
