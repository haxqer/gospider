package parser

import (
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/pkg/setting"
	"regexp"
)

var channelRe = regexp.MustCompile(`<a[^>]*href="(//www\.mgtv\.com/\w+/\d+/\d+\.html)"[^>]*>([^<]+)</a>`)
var episodeRe = regexp.MustCompile(`//www\.mgtv\.com/\w+/\d+/(\d+)\.html`)

func ParseChannel(contents []byte, channelID string) engine.ParseResult {
	result := engine.ParseResult{}

	matches := channelRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		matchesEpisode := episodeRe.FindSubmatch(m[1])
		episodeID := string(matchesEpisode[1])
		url := GenEpisodeAPIURLByEpisodeID(episodeID, 1)
		if !setting.ServerSetting.IsFull {
			url = GenEpisodeAPIURLByEpisodeIDNew(episodeID, 1)
		}
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseEpisode(c, url, channelID)
			},
		})
	}

	return result
}
