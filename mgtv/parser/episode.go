package parser

import (
	"fmt"
	"git.trac.cn/nv/spider/engine"
	"regexp"
)

var episodeApiRe = regexp.MustCompile(`^jQuery\d+_\d+\((.*)\)$`)

func ParseEpisode(contents []byte) engine.ParseResult {
	matches := episodeApiRe.FindSubmatch(contents)
	jsonStr := matches[1]

	fmt.Printf("%s \n", jsonStr)

	result := engine.ParseResult{}

	return result
}
