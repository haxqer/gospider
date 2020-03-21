package model

type Episode struct {
	DramaID     string `json:"clip_id"`
	ContentType string `json:"contentType"`
	Image       string `json:"img"`
	IsIntact    string `json:"isIntact"`
	IsNew       string `json:"isnew"`
	IsVIP       string `json:"isvip"`
	PlayCounter string `json:"playcnt"`
	Title1      string `json:"t1"`
	Title2      string `json:"t2"`
	Title3      string `json:"t3"`
	Title4      string `json:"t4"`
	Duration    string `json:"time"`
	TS          string `json:"ts"`
	URL         string `json:"url"`
	EpisodeID   string `json:"video_id"`
}
