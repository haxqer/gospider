package parser

import (
	"fmt"
	"git.trac.cn/nv/spider/engine"
	"math/rand"
	"regexp"
	"time"
)

var channelRe = regexp.MustCompile(`<a[^>]*href="(//www\.mgtv\.com/\w+/\d+/\d+\.html)"[^>]*>([^<]+)</a>`)
var episodeRe = regexp.MustCompile(`//www\.mgtv\.com/\w+/\d+/(\d+)\.html`)

func ParseChannel(contents []byte) engine.ParseResult {
	matches := channelRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "drama "+string(m[2]))

		matchesEpisode := episodeRe.FindSubmatch(m[1])
		videoID := matchesEpisode[1]
		jpRand := rand.Int63n(8030056088838044) + 1030056088838044
		nowTS := time.Now().UnixNano() / int64(time.Millisecond)
		jqTS := nowTS - rand.Int63n(400) + 100
		page := 1

		result.Requests = append(result.Requests, engine.Request{
			//Url:        "https:" + string(m[1]),
			Url: fmt.Sprintf("https://pcweb.api.mgtv.com/episode/list?video_id=%s"+
				"&page=%d&size=25"+
				"&cxid=&version=5.5.35&callback=jQuery1820%d_%d&_support=10000000&_=%d",
				videoID, page, jpRand, nowTS, jqTS),
			ParserFunc: engine.NilParser,
		})

	}

	return result
}
