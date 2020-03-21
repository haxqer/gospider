package parser

import (
	"git.trac.cn/nv/spider/engine"
	"regexp"
)

var channelListRe = regexp.MustCompile(`<li><a[^>]*href="(/-------------\.html\?channelId=\d+)">([^<]+)</a></li>`)

func ParseChannelList(contents []byte) engine.ParseResult {
	matches := channelListRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://list.mgtv.com" + string(m[1]),
			ParserFunc: engine.NilParser,
		})

	}

	return result
}
