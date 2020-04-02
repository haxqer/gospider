package main

import (
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/fetcher"
	"git.trac.cn/nv/spider/mgtv/parser"
	"git.trac.cn/nv/spider/model"
	"git.trac.cn/nv/spider/persist"
	"git.trac.cn/nv/spider/pkg/logging"
	"git.trac.cn/nv/spider/pkg/setting"
	"git.trac.cn/nv/spider/scheduler"
)

func init() {
	setting.Setup()
	logging.Setup()
	fetcher.Setup()
	model.Setup()
	engine.Setup()
}

func main() {
	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: setting.ServerSetting.WorkerCount,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://list.mgtv.com/-------------.html?channelId=1",
		ParserFunc: parser.ParseChannelList,
	})

}
