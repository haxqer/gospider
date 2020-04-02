package parser

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
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

func DurationUnmarshalText(s string) (time.Duration, error) {
	var t time.Duration
	s = strings.TrimSpace(s)
	if s == "" || strings.ToLower(s) == "undefined" {
		return 0, fmt.Errorf("invalid duration: %s", s)
	}
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid duration: %s", s)
	}
	if i := strings.IndexByte(parts[1], '.'); i > 0 {
		ms, err := strconv.ParseInt(parts[1][i+1:], 10, 32)
		if err != nil || ms < 0 || ms > 999 {
			return 0, fmt.Errorf("invalid duration: %s", s)
		}
		parts[1] = parts[1][:i]
		t += time.Duration(ms) * time.Millisecond
	}
	f := time.Second
	for i := 1; i >= 0; i-- {
		n, err := strconv.ParseInt(parts[i], 10, 32)
		if err != nil || n < 0 {
			return 0, fmt.Errorf("invalid duration: %s", s)
		}
		t += time.Duration(n) * f
		f *= 60
	}

	return t, nil
}

func ParseMgtvTime(s string) (time.Time, error) {
	layout := "2006-01-02 15:04:05.0"
	t, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

var playCounterRe = regexp.MustCompile(`^([0-9.]+)(.*)$`)

func ParseMgtvPlayCounter(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0, fmt.Errorf("parse playCounter error(is empty): %s", s)
	}
	m := playCounterRe.FindSubmatch([]byte(s))
	if len(m) != 3 {
		return 0, fmt.Errorf("parse playCounter error: %s", s)
	}
	number, err := strconv.ParseFloat(string(m[1]), 64)
	if err != nil {
		return 0, err
	}

	m2 := strings.TrimSpace(string(m[2]))
	if m2 == "万" {
		number = number * 10000
	}

	if m2 == "亿" {
		number = number * 100000000
	}
	//

	return int64(number), nil
}
