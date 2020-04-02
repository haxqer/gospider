package engine

import (
	hasher2 "git.trac.cn/nv/spider/pkg/hasher"
	"git.trac.cn/nv/spider/pkg/setting"
	"github.com/patrickmn/go-cache"
	"time"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
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

	tickTrigger(e, seeds)
	trigger(e, seeds)

	for {
		result := <-results
		for _, item := range result.Items {
			go func(i Item) {
				e.ItemChan <- i
			}(item)
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var seedTrigger <-chan time.Time
func tickTrigger(e *ConcurrentEngine, seeds []Request)  {
	go func() {
		for {
			<-seedTrigger
			trigger(e, seeds)
		}
	}()
}

func trigger(e *ConcurrentEngine, seeds []Request)  {
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}
}

var visitedUrls *cache.Cache

func Setup() {
	seedTrigger = time.Tick(1 * time.Minute)
	visitedUrls = cache.New(setting.ServerSetting.UrlExpire, setting.ServerSetting.UrlExpire + 5*time.Minute)
}

func isDuplicate(url string) bool {
	hasher := hasher2.GetMD5Hash(url)
	if _, found := visitedUrls.Get(hasher); found {
		return true
	}
	visitedUrls.SetDefault(hasher, true)
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
