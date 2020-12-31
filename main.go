package main

import (
	"fmt"
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/fetcher"
	"git.trac.cn/nv/spider/mgtv"
	"git.trac.cn/nv/spider/mgtv/parser"
	"git.trac.cn/nv/spider/persist"
	"git.trac.cn/nv/spider/pkg/logging"
	"git.trac.cn/nv/spider/pkg/setting"
	"git.trac.cn/nv/spider/scheduler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	logging.Setup()
	fetcher.Setup()
	engine.Setup()
	persist.Setup()
}

func main() {
	go func() {
		metricsEndPoint := fmt.Sprintf(":%d", setting.ServerSetting.MetricsPort)
		http.Handle("/metrics", promhttp.Handler())
		log.Printf("[info] start metrics server of prometheus listening %s", metricsEndPoint)
		http.ListenAndServe(metricsEndPoint, nil)
	}()

	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: setting.ServerSetting.WorkerCount,
		ItemChan:    itemChan,
	}
	if setting.ServerSetting.IsFull {
		e.Run(engine.Request{
			Url:        "http://list.mgtv.com/-------------.html?channelId=1",
			ParserFunc: parser.ParseChannelList,
		})
	} else {
		e.Run(
			genChannelRequest("1"),
			genChannelRequest("2"),
			genChannelRequest("3"),
			genChannelRequest("10"),
			genChannelRequest("50"),
		)
	}
}

func genChannelRequest(channelId string) engine.Request {
	return engine.Request{
		Url: "https://pianku.api.mgtv.com/rider/list/pcweb/v3?platform=pcweb&channelId=" + channelId +
			"&pn=1&pc=80&hudong=1&_support=10000000&kind=a1&area=a1&sort=c2&callback=" + mgtv.GenJsonp(),
		ParserFunc: func(c []byte) engine.ParseResult {
			return parser.ParseChannel(c, channelId)
		},
	}
}

func genUrlRequest(channelId string, url string) engine.Request {
	return engine.Request{
		Url: url,
		ParserFunc: func(c []byte) engine.ParseResult {
			return parser.ParseChannel(c, channelId)
		},
	}
}
