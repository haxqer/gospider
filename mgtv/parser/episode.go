package parser

import (
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/model"
	"github.com/pquerna/ffjson/ffjson"
	"log"
	"net/http"
	"regexp"
)

var episodeApiRe = regexp.MustCompile(`^jQuery\d+_\d+\((.*)\)$`)

type episodeAPI struct {
	Code int             `json:"code"`
	Data *episodeAPIData `json:"data"`
}

type episodeAPIData struct {
	Info        *dramaInfo      `json:"info"`
	Total       int             `json:"total"`
	Count       int             `json:"count"`
	TotalPage   int             `json:"total_page"`
	CurrentPage int             `json:"current_page"`
	EpisodeList []model.Episode `json:"list"`
}

type dramaInfo struct {
	Title string `json:"title"`
	Type  string `json:"type"`
	IsVIP string `json:"isvip"`
	Desc  string `json:"desc"`
}

func ParseEpisode(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	matches := episodeApiRe.FindSubmatch(contents)
	if len(matches) < 2 {
		return result
	}

	jsonStr := matches[1]

	var episodeResult episodeAPI
	err := ffjson.Unmarshal(jsonStr, &episodeResult)
	if err != nil {
		log.Printf("ffjson.Unmarshal: error "+"jsonStr %s: %v", jsonStr, err)
		return result
	}
	if episodeResult.Code != http.StatusOK {
		return result
	}

	data := episodeResult.Data
	if len(data.EpisodeList) == 0 {
		return result
	}

	tempEpisodeID := data.EpisodeList[0].EpisodeID

	if data.TotalPage > data.CurrentPage {
		for i := data.CurrentPage + 1; i <= data.TotalPage; i++ {
			result.Requests = append(result.Requests, engine.Request{
				Url:        GenEpisodeAPIURLByEpisodeID(tempEpisodeID, i),
				ParserFunc: ParseEpisode,
			})
		}
	}

	for _, episode := range data.EpisodeList {
		result.Items = append(result.Items, []interface{}{episode})
	}

	return result
}
