package model

import (
	"encoding/json"
)

type Mgtv struct {
	EpisodeId   int    `gorm:"primary_key" json:"episode_id"`
	ChannelId   int    `json:"channel_id"`
	DramaId     int    `json:"drama_id"`
	DramaTitle  string `json:"drama_title"`
	Title1      string `json:"title1"`
	Title2      string `json:"title2"`
	Title3      string `json:"title3"`
	Title4      string `json:"title4"`
	EpisodeUrl  string `json:"episode_url"`
	Duration    int    `json:"duration"`
	ContentType string `json:"content_type"`
	Image       string `json:"image"`
	IsIntact    string `json:"is_intact"`
	IsNew       string `json:"is_new"`
	IsVip       string `json:"is_vip"`
	PlayCounter string `json:"play_counter"`
	Ts          string `json:"ts"`
	NextId      string `json:"next_id"`
	SrcClipId   string `json:"src_clip_id"`
}

func FromJsonObj(o interface{}) (Mgtv, error) {
	var episode Mgtv
	s, err := json.Marshal(o)
	if err != nil {
		return episode, err
	}
	err = json.Unmarshal(s, &episode)
	return episode, err
}

func InsertOnDuplicate(mgtv *Mgtv) error {
	err := db.Save(mgtv).Error
	if err != nil {
		return err
	}
	return nil
}
