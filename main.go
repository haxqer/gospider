package main

import (
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/mgtv/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "https://list.mgtv.com/1/a1-a1--------c1-25---.html?channelId=1",
		ParserFunc: parser.ParseChannelList,
	})

}