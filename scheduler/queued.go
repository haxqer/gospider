package scheduler

import "git.trac.cn/nv/spider/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkerReady(worker chan engine.Request) {
	q.workerChan <- worker
}

func (q *QueuedScheduler) Run() {
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
			case request := <-q.requestChan:
				requestQueue = append(requestQueue, request)
			case worker := <-q.workerChan:
				workerQueue = append(workerQueue, worker)
			case activeWorker <- activeRequest:
				workerQueue = workerQueue[1:]
				requestQueue = requestQueue[1:]
			}
		}
	}()
}
