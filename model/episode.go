package model

import "encoding/json"

type Episode struct {
	ChannelID   string
	DramaID     string
	DramaTitle  string
	EpisodeID   string
	Title1      string
	Title2      string
	Title3      string
	Title4      string
	EpisodeURL  string
	Duration    string
	ContentType string
	Image       string
	IsIntact    string
	IsNew       string
	IsVIP       string
	PlayCounter string
	TS          string
	NextID      string
	SrcClipID   string
}

func FromJsonObj(o interface{}) (Episode, error) {
	var episode Episode
	s, err := json.Marshal(o)
	if err != nil {
		return episode, err
	}
	err = json.Unmarshal(s, &episode)
	return episode, err
}
