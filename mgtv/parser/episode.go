package parser

import (
	"fmt"
	"git.trac.cn/nv/spider/engine"
	"git.trac.cn/nv/spider/model"
	"git.trac.cn/nv/spider/pkg/logging"
	"git.trac.cn/nv/spider/pkg/setting"
	"github.com/pquerna/ffjson/ffjson"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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

func ParseEpisode(contents []byte, url string, channelID string) engine.ParseResult {
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

	if data.TotalPage > data.CurrentPage && setting.ServerSetting.IsFull {
		for i := data.CurrentPage + 1; i <= data.TotalPage; i++ {
			result.Requests = append(result.Requests, engine.Request{
				Url: GenEpisodeAPIURLByEpisodeID(tempEpisodeID, i),
				ParserFunc: func(c []byte) engine.ParseResult {
					return ParseEpisode(c, url, channelID)
				},
			})
		}
	}

	for _, m := range data.EpisodeList {

		intChannelID, err := strconv.Atoi(strings.TrimSpace(channelID))
		if err != nil {
			logging.Error(fmt.Sprintf("channelID %s parse error: %+v", channelID, err))
			continue
		}

		intDramaID, err := strconv.Atoi(strings.TrimSpace(m.DramaID))
		if err != nil {
			logging.Error(fmt.Sprintf("DramaID %s parse error: %+v", m.DramaID, err))
			continue
		}

		intEpisodeID, err := strconv.Atoi(strings.TrimSpace(m.EpisodeID))
		if err != nil {
			logging.Error(fmt.Sprintf("EpisodeID %s parse error: %+v", m.EpisodeID, err))
			continue
		}

		timeDuration, err := DurationUnmarshalText(m.Duration)
		if err != nil {
			logging.Error(fmt.Sprintf("Duration %s parse error: %+v", m.Duration, err))
			continue
		}
		intDuration := int32(timeDuration.Seconds())

		int64PlayCounter, _ := ParseMgtvPlayCounter(m.PlayCounter)
		//if err != nil {
		//	logging.Error(fmt.Sprintf("PlayCounter %s parse error: %+v", m.PlayCounter, err))
		//	continue
		//}
		if int64PlayCounter == 0 && intDuration >= 600 {
			int64PlayCounter = 100000
		}

		item := model.Mgtv{
			ChannelId:   int32(intChannelID),
			DramaId:     int32(intDramaID),
			DramaTitle:  data.Info.Title,
			EpisodeId:   int32(intEpisodeID),
			Title1:      m.Title1,
			Title2:      m.Title2,
			Title3:      m.Title3,
			Title4:      m.Title4,
			EpisodeUrl:  "http://www.mgtv.com" + m.URL,
			Duration:    intDuration,
			ContentType: m.ContentType,
			Image:       m.Image,
			IsIntact:    m.IsIntact,
			IsNew:       m.IsNew,
			IsVip:       m.IsVIP,
			PlayCounter: int64PlayCounter,
			Ts:          m.TS,
			NextId:      m.NextID,
			SrcClipId:   m.SrcClipID,
		}
		result.Items = append(result.Items, item)
	}

	return result
}

func EpisodeFunc(channelID string, url string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseEpisode(c, url, channelID)
	}
}
