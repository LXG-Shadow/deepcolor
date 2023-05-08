package backup

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

func wait() {
	limiter := rate.NewLimiter(3, 5)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	for i := 0; ; i++ {
		fmt.Printf("%03d %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return
		}
	}
}

func allow() {
	limiter := rate.NewLimiter(3, 5)
	for i := 0; i < 50; i++ {
		if limiter.Allow() {
			fmt.Printf("%03d Ok  %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		} else {
			fmt.Printf("%03d Err %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func TestLimiter(t *testing.T) {
	allow()
}
