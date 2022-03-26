package deepcolor

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type Engine struct {
	lock          sync.RWMutex
	waitGroup     sync.WaitGroup
	limiter       *rate.Limiter
	waitChan      chan int
	requestQueue  *QueueChannel
	maxConnection int
	context       context.Context
	requestFunc   RequestFunc
	ReqHandlers   []RequestHandler
	RespHandlers  []ResponseHandler
}

func NewEngine(maxConn int) *Engine {
	e := &Engine{
		requestFunc:   Get,
		requestQueue:  NewQueueChannel(maxConn),
		maxConnection: maxConn,
		context:       context.Background(),
		limiter:       rate.NewLimiter(rate.Every(time.Millisecond*10), 1),
		ReqHandlers:   make([]RequestHandler, 0),
		RespHandlers:  make([]ResponseHandler, 0),
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
		fmt.Println("Worker start for", t.(Tentacle).Url)
		e.FetchTentacle(t.(Tentacle))
		e.waitGroup.Done()
		fmt.Println("Worker finish for", t.(Tentacle).Url)
	}
}

func (e *Engine) FetchTentacle(tentacle Tentacle) *TentacleResult {
	err := e.limiter.WaitN(e.context, 1)
	if err != nil {
		return nil
	}
	result, _ := Fetch(tentacle, e.requestFunc, e.ReqHandlers, e.RespHandlers)
	return result
}

func (e *Engine) Fetch(uri string) *TentacleResult {
	return e.FetchTentacle(TentacleHTML(uri, "utf-8"))
}
func (e *Engine) FetchTentacleAsync(tentacle Tentacle) {
	e.waitGroup.Add(1)
	e.requestQueue.Push(tentacle)

	//go func() {
	//	e.FetchTentacle(tentacle)
	//	defer e.waitGroup.Done()
	//}()
}

func (e *Engine) FetchAsync(uri string) {
	e.FetchTentacleAsync(TentacleHTML(uri, "utf-8"))
	//go func() {
	//	e.Fetch(uri)
	//	defer e.waitGroup.Done()
	//}()
}

func (e *Engine) WaitUntilFinish() {
	e.waitGroup.Wait()
}

func (e *Engine) SetPeriod(duration time.Duration) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.limiter.SetLimit(rate.Every(duration))
}

func (e *Engine) SetBurst(burst int) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.limiter.SetBurst(burst)
}

//func (e *Engine) SetMaxConnection(conn int) {
//	e.lock.Lock()
//	defer e.lock.Unlock()
//	e.waitChan = make(chan int, conn)
//	for i := 1; i <= conn; i++ {
//		e.waitChan <- i
//	}
//}

func (e *Engine) SetRequestFunc(requestFunc RequestFunc) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.requestFunc = requestFunc
}

func (e *Engine) OnRequest(handlers ...RequestHandler) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.ReqHandlers = append(e.ReqHandlers, handlers...)
}

func (e *Engine) OnResponse(handlers ...ResponseHandler) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.RespHandlers = append(e.RespHandlers, handlers...)
}
