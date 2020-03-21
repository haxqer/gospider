package engine

import (
	"git.trac.cn/nv/spider/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)

		bytes, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error "+"fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(bytes)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {

			log.Printf("Got item: %s", item)
		}

	}
}
