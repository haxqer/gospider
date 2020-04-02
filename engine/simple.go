package engine

import (
	"fmt"
	"git.trac.cn/nv/spider/fetcher"
	"git.trac.cn/nv/spider/pkg/logging"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {

			log.Printf("Got item: %s", item)
		}

	}
}

func worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		logging.Error(fmt.Sprintf("Fetcher: error: %v "+"fetching url %s", err, r.Url))
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
