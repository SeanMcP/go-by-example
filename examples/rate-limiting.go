// rate-limiting.go
package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// time.Tick returns a receive-only channel
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request       ", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	// first three requests are allowed immediately
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// after that, we allow a request every 200ms
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("bursty request", req, time.Now())
	}
}
