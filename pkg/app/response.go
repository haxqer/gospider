package app

import (
	"git.trac.cn/nv/spider/model"
	"git.trac.cn/nv/spider/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type SpiderResponse struct {
	Code   int          `json:"code"`
	Msg    string       `json:"msg"`
	Length int          `json:"length"`
	Data   []model.Mgtv `json:"data"`
	Error  string       `json:"error"`
}

func (g *Gin) EpisodeResponse(httpCode, errCode int, data []model.Mgtv, err error) {

	errStr := ""
	if err != nil {
		errStr = err.Error()
	}
	g.C.JSON(httpCode, SpiderResponse{
		Code:   errCode,
		Msg:    e.GetMsg(errCode),
		Length: len(data),
		Data:   data,
		Error: errStr,
	})
	return
}
