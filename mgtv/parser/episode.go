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
	Info        *dramaInfo    `json:"info"`
	Total       int           `json:"total"`
	Count       int           `json:"count"`
	TotalPage   int           `json:"total_page"`
	CurrentPage int           `json:"current_page"`
	EpisodeList []mgtvEpisode `json:"list"`
}

type dramaInfo struct {
	Title string `json:"title"`
	Type  string `json:"type"`
	IsVIP string `json:"isvip"`
	Desc  string `json:"desc"`
}

type mgtvEpisode struct {
	DramaID     string `json:"clip_id"`
	EpisodeID   string `json:"video_id"`
	Title1      string `json:"t1"`
	Title2      string `json:"t2"`
	Title3      string `json:"t3"`
	Title4      string `json:"t4"`
	URL         string `json:"url"`
	Duration    string `json:"time"`
	ContentType string `json:"contentType"`
	Image       string `json:"img"`
	IsIntact    string `json:"isIntact"`
	IsNew       string `json:"isnew"`
	IsVIP       string `json:"isvip"`
	PlayCounter string `json:"playcnt"`
	TS          string `json:"ts"`
	NextID      string `json:"next_id"`
	SrcClipID   string `json:"src_clip_id"`
}

func ParseEpisode(contents []byte, channelID string) engine.ParseResult {
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
				Url: GenEpisodeAPIURLByEpisodeID(tempEpisodeID, i),
				ParserFunc: func(c []byte) engine.ParseResult {
					return ParseEpisode(c, channelID)
				},
			})
		}
	}

	for _, m := range data.EpisodeList {
		episode := model.Episode{
			ChannelID:   channelID,
			DramaID:     m.DramaID,
			DramaTitle:  data.Info.Title,
			EpisodeID:   m.EpisodeID,
			Title1:      m.Title1,
			Title2:      m.Title2,
			Title3:      m.Title3,
			Title4:      m.Title4,
			URL:         "https://www.mgtv.com/" + m.URL,
			Duration:    m.Duration,
			ContentType: m.ContentType,
			Image:       m.Image,
			IsIntact:    m.IsIntact,
			IsNew:       m.IsNew,
			IsVIP:       m.IsVIP,
			PlayCounter: m.PlayCounter,
			TS:          m.TS,
			NextID:      m.NextID,
			SrcClipID:   m.SrcClipID,
		}
		result.Items = append(result.Items, episode)
	}

	return result
}
