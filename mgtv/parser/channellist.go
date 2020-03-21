package parser

import (
	"git.trac.cn/nv/spider/engine"
	"regexp"
)

//var channelListRe = regexp.MustCompile(`<li><a[^>]*href="(/-------------\.html\?channelId=\d+)">([^<]+)</a></li>`)
var channelListRe = regexp.MustCompile(`<li><a[^>]*href="([^>]+\.html\?channelId=)(\d+)"`)

func ParseChannelList(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}

	matches := channelListRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		channelID := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: "https://list.mgtv.com" + string(m[1]) + channelID,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseChannel(c, channelID)
			},
		})
	}

	return result
}
