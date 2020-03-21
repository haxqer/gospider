package model

type Episode struct {
	ChannelID   string
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
}
