package engine

import (
	"github.com/pquerna/ffjson/ffjson"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	results := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), results, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	episodeCount := 0
	for {
		result := <-results
		for _, item := range result.Items {
			v, _ := ffjson.Marshal(item)
			log.Printf("Got item #%d: %s", episodeCount, v)
			episodeCount++
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}

func createWorker(requests chan Request, results chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(requests)
			request := <-requests
			result, err := worker(request)
			if err != nil {
				continue
			}
			results <- result
		}
	}()
}
