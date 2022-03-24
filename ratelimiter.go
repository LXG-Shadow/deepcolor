package deepcolor

import (
	"sync"
)

type WaitGroup struct {
	internal sync.WaitGroup
	waitChan chan int
}

func (wg *WaitGroup) SetMaxConnection(c int) {
	wg.waitChan = make(chan int, c)
	for i := 1; i <= 5; i++ {
		wg.waitChan <- i
	}
}

func (wg *WaitGroup) Add(delta int) {
	for i := 0; i < delta; i++ {
		<-wg.waitChan
	}
	wg.internal.Add(delta)
}

func (wg *WaitGroup) Done() {
	wg.DoneN(1)
}

func (wg *WaitGroup) DoneN(delta int) {
	for i := 0; i < delta; i++ {
		wg.waitChan <- i
	}
	wg.internal.Add(-delta)
}

func (wg *WaitGroup) Wait() {
	wg.internal.Wait()
}
