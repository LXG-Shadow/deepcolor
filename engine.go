package deepcolor

import (
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type ProcessFunc func()

type Engine struct {
	lock          sync.RWMutex
	waitGroup     sync.WaitGroup
	limiter       *rate.Limiter
	waitChan      chan int
	requestQueue  *QueueChannel
	maxConnection int
	context       context.Context
}

func NewEngine(maxConn int) *Engine {
	e := &Engine{
		requestQueue:  NewQueueChannel(maxConn),
		maxConnection: maxConn,
		context:       context.Background(),
		limiter:       rate.NewLimiter(rate.Every(time.Millisecond*10), 1),
	}
	e.makeAsyncWorkers()
	return e
}

func (e *Engine) makeAsyncWorkers() {
	for i := 0; i < e.maxConnection; i++ {
		go e.newAsyncWorker()
	}
}

func (e *Engine) newAsyncWorker() {
	for t := range e.requestQueue.Chan {
		err := e.limiter.WaitN(e.context, 1)
		if err != nil {
			continue
		}
		(t.(ProcessFunc))()
		e.waitGroup.Done()
	}
}

func (e *Engine) Add(dp *Deepcolor) {
	var req = dp.Requester.Next(nil)
	for req != nil {
		r := req
		e.waitGroup.Add(1)
		e.requestQueue.Chan <- ProcessFunc(func() {
			for _, fun := range dp.ReqHandler {
				if !fun(r) {
					return
				}
			}
			resp := dp.ReqFunc(r)
			for _, fun := range dp.RespHandler {
				if !fun(resp) {
					return
				}
			}
			for _, t := range dp.Tentacles {
				err := t.Initialize(resp)
				if err != nil {
					continue
				}
				for _, h := range t.Handlers {
					h(t)
				}
			}
		})
		req = dp.Requester.Next(nil)
	}
	return
}

func (e *Engine) WaitUntilFinish() {
	e.waitGroup.Wait()
}

func (e *Engine) SetPeriod(duration time.Duration) {
	e.limiter.SetLimit(rate.Every(duration))
}

func (e *Engine) SetBurst(burst int) {
	e.limiter.SetBurst(burst)
}
