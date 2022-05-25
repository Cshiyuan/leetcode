package awesome

import (
	"fmt"
	"golang.org/x/time/rate"
	"sync/atomic"
	"time"
)

func test() {

	var succCount, failCount int64
	limit := rate.Every(100 * time.Millisecond)
	burst := 1
	limiter := rate.NewLimiter(limit, burst)
	start := time.Now()
	for i := 0; i < 5000; i++ {
		go func() {
			for {
				if limiter.Allow() {
					atomic.AddInt64(&succCount, 1)
				} else {
					atomic.AddInt64(&failCount, 1)
				}
			}
		}()
	}
	time.Sleep(2 * time.Second)
	elapsed := time.Since(start)
	fmt.Println("elapsed=", elapsed, "succCount=", atomic.LoadInt64(&succCount), "failCount=", atomic.LoadInt64(&failCount))

}
