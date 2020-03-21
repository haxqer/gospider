package parser

import (
	"git.trac.cn/nv/spider/engine"
	"regexp"
)

var channelRe = regexp.MustCompile(`<a[^>]*href="(//www\.mgtv\.com/\w+/\d+/\d+\.html)"[^>]*>([^<]+)</a>`)
var episodeRe = regexp.MustCompile(`//www\.mgtv\.com/\w+/\d+/(\d+)\.html`)

func ParseChannel(contents []byte, channelID string) engine.ParseResult {
	result := engine.ParseResult{}

	matches := channelRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Items = append(result.Items, "drama "+string(m[2]))
		matchesEpisode := episodeRe.FindSubmatch(m[1])
		episodeID := string(matchesEpisode[1])

		result.Requests = append(result.Requests, engine.Request{
			Url: GenEpisodeAPIURLByEpisodeID(episodeID, 1),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseEpisode(c, channelID)
			},
		})
	}

	return result
}
