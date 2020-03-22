package main

import (
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/mgtv/parser"
	"git.trac.cn/nv/spider/persist"
	"git.trac.cn/nv/spider/scheduler"
)

func main() {
	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        "https://list.mgtv.com/-------------.html?channelId=1",
		ParserFunc: parser.ParseChannelList,
	})

}
