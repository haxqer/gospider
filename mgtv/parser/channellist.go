package parser

import (
	"git.trac.cn/nv/spider/engine"
	"regexp"
)

var channelListRe = regexp.MustCompile(`<li><a[^>]*href="(/-------------\.html\?channelId=\d+)">([^<]+)</a></li>`)

//var nextChannelPageRe = regexp.MustCompile(`<li><a href="([^>]+\.html\?channelId=\d+)" class="next turn" title="下一页">`)

func ParseChannelList(contents []byte) engine.ParseResult {
	matches := channelListRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "Channel "+string(m[2])+" page 1")
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://list.mgtv.com" + string(m[1]),
			ParserFunc: ParseChannel,
		})

	}
	//
	//nextChannelPageMatches := nextChannelPageRe.FindSubmatch(contents)
	//if len(nextChannelPageMatches) > 2 {
	//	result.Items = append(result.Items, "Channel "+string(m[2]))
	//}

	return result
}
