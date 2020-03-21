package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	requests := make(chan Request)
	results := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(requests)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(requests, results)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-results
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(requests chan Request, results chan ParseResult) {
	go func() {
		for {
			request := <-requests
			result, err := worker(request)
			if err != nil {
				continue
			}
			results <- result
		}
	}()
}
