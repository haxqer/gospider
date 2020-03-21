package main

import (
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/mgtv/parser"
	"git.trac.cn/nv/spider/scheduler"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "https://list.mgtv.com/1/a1-a1--------c1-25---.html?channelId=1",
		ParserFunc: parser.ParseChannelList,
	})

}
