package parser

import (
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/pkg/setting"
	"github.com/pquerna/ffjson/ffjson"
	"log"
	"net/http"
	"regexp"
)

var channelRe = regexp.MustCompile(`href="(//www\.mgtv\.com/\w+/\d+/\d+\.html)"`)
var episodeRe = regexp.MustCompile(`//www\.mgtv\.com/\w+/\d+/(\d+)\.html`)

var channelListApiRe = regexp.MustCompile(`^jsonp_\d+_\d+\(([^)]+)\)$`)

type channelListAPI struct {
	Code int             `json:"code"`
	Data *channelListAPIData `json:"data"`
}

type channelListAPIData struct {
	Total       int           `json:"totalHits"`
	ChannelList []mgtvDrama `json:"hitDocs"`
}

type mgtvDrama struct {
	EpisodeId string `json:"playPartId"`
}

func ParseChannel(contents []byte, channelID string) engine.ParseResult {
	result := engine.ParseResult{}

	matches := channelListApiRe.FindSubmatch(contents)
	if len(matches) < 2 {
		return result
	}

	jsonStr := matches[1]

	var channelResult channelListAPI
	err := ffjson.Unmarshal(jsonStr, &channelResult)
	if err != nil {
		log.Printf("ffjson.Unmarshal: error "+"jsonStr %s: %v", jsonStr, err)
		return result
	}
	if channelResult.Code != http.StatusOK {
		return result
	}

	data := channelResult.Data
	if len(data.ChannelList) == 0 {
		return result
	}
	channelList := data.ChannelList

	for _, c := range channelList {
		episodeID := c.EpisodeId
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
