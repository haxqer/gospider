package routers

import (
	"fmt"
	"git.trac.cn/nv/spider/fetcher"
	"git.trac.cn/nv/spider/mgtv/parser"
	"git.trac.cn/nv/spider/model"
	"git.trac.cn/nv/spider/persist"
	"git.trac.cn/nv/spider/pkg/app"
	"git.trac.cn/nv/spider/pkg/e"
	"git.trac.cn/nv/spider/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

var episodeRe = regexp.MustCompile(`https?://www\.mgtv\.com/\w+/\d+/(\d+)\.html`)

// episode
// @tags Spider
// @Summary Insert episode from URL of MGTV
// @Description insert episode
// @ID insert episode
// @Produce json
// @Param name body app.SpiderRequest true "Name"
// @Success 200 {object} app.SpiderResponse
// @Failure 400 {object} app.SpiderResponse
// @Router /v1/spider [post]
func Episode(c *gin.Context) {
	appG := app.Gin{C: c}

	sr := &app.SpiderRequest{}
	err := c.ShouldBindJSON(sr)
	if err != nil {
		appG.EpisodeResponse(http.StatusBadRequest, e.InvalidParams, nil, err)
		return
	}

	url := genUrl(sr)

	body, err := fetcher.Fetch(url)
	if err != nil {
		logging.Error(fmt.Sprintf("Fetcher: error: %v "+" fetching url %s", err, sr.Url))
		appG.EpisodeResponse(http.StatusBadRequest, e.ERROR, nil, err)
		return
	}

	parseResult := parser.ParseEpisode(body, url, strconv.Itoa(sr.ChannelId))

	err = save(parseResult.Items)
	if err != nil {
		logging.Error(err)
		appG.EpisodeResponse(http.StatusInternalServerError, e.ERROR, nil, err)
		return
	}

	appG.EpisodeResponse(http.StatusOK, e.SUCCESS, parseResult.Items, nil)
}

func save(items []model.Mgtv) error {
	var wg sync.WaitGroup
	ce := make(chan error)
	defer close(ce)
	for _,m := range items {
		wg.Add(1)

		go worker(m, &wg, ce)
	}
	wg.Wait()

	select {
	case err := <-ce:
		return err
	default:
		return nil
	}
}

func worker(m model.Mgtv, wg *sync.WaitGroup, ce chan error) {
	defer wg.Done()
	err := persist.RpcCall(&m)
	if err != nil {
		ce <- err
	}
}

func genUrl(sr *app.SpiderRequest) string {
	matchesEpisode := episodeRe.FindSubmatch([]byte(sr.Url))
	episodeID := string(matchesEpisode[1])
	url := parser.GenEpisodeAPIURLByEpisodeID(episodeID, 1)
	return url
}
