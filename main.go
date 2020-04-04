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
	persist.Setup()
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
	if setting.ServerSetting.IsFull {
		e.Run(engine.Request{
			Url:        "http://list.mgtv.com/-------------.html?channelId=1",
			ParserFunc: parser.ParseChannelList,
		})
	}else{
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
		Url:        "http://list.mgtv.com/-------------.html?channelId=" + channelId,
		ParserFunc: func(c []byte) engine.ParseResult {
			return parser.ParseChannel(c, channelId)
		},
	}
}