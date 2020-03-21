package main

import (
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/mgtv/parser"
	"git.trac.cn/nv/spider/scheduler"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "https://list.mgtv.com/-------------.html?channelId=1",
		ParserFunc: parser.ParseChannelList,
	})

}
