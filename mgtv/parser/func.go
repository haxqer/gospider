package parser

import (
	"fmt"
	"math/rand"
	"time"
)

const pageSize = 50

func GenEpisodeAPIURLByEpisodeID(episodeID string, page int) string {
	jpRand := rand.Int63n(8030056088838044) + 1030056088838044
	nowTS := time.Now().UnixNano() / int64(time.Millisecond)
	jqTS := nowTS - rand.Int63n(400) + 100
	return fmt.Sprintf("https://pcweb.api.mgtv.com/episode/list?video_id=%s"+
		"&page=%d&size=%d"+
		"&cxid=&version=5.5.35&callback=jQuery1820%d_%d&_support=10000000&_=%d",
		episodeID, page, pageSize, jpRand, nowTS, jqTS)
}
