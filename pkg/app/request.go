package app

type SpiderRequest struct {
	ChannelId int    `json:"channel_id" binding:"required,gt=0"`
	Url       string `json:"url" binding:"required,mgtv_url"`
}
