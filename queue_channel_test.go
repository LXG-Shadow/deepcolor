package deepcolor

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestQueueChannel(t *testing.T) {
	qc := NewQueueChannel(128)
	var wg sync.WaitGroup
	wg.Add(1024 * 1234)
	go func() {
		for x := range qc.Chan {
			if x == nil {
				t.Errorf("Fail")
			}
			wg.Done()
		}
	}()

	for i := 0; i < 1024; i++ {
		go func() {
			for j := 0; j < 1234; j++ {
				qc.Push(j)
			}
		}()
	}
	fmt.Println("Finish add")
	wg.Wait()
}

func TestQueueChannel1(t *testing.T) {
	qc := NewQueueChannel(128)
	var wg sync.WaitGroup
	wg.Add(1024)
	for i := 0; i < 10; i++ {
		ii := i
		go func() {
			for x := range qc.Chan {
				fmt.Printf("%d done in worker %d\n", x, ii)
				wg.Done()
				time.Sleep(time.Second * time.Duration(rand.Intn(3)+2))
			}
		}()
	}

	for j := 0; j < 1024; j++ {
		qc.Push(j)
	}
	fmt.Println("Finish add")
	wg.Wait()
}
