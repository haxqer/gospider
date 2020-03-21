package parser

import (
	"git.trac.cn/nv/spider/engine"
	"regexp"
)

var channelRe = regexp.MustCompile(`<a[^>]*href="(//www\.mgtv\.com/\w+/\d+/\d+\.html)"[^>]*>([^<]+)</a>`)

func ParseChannel(contents []byte) engine.ParseResult {
	matches := channelRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "drama "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https:" + string(m[1]),
			ParserFunc: engine.NilParser,
		})

	}

	return result
}
