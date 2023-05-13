// atomic-counters.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64 // counter

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1) // hey waitgroup, we are starting a goroutine

		go func() {
			for c := 0; c < 1000; c++ {
				// method guarantees that goroutine increments the counter
				atomic.AddUint64(&ops, 1)
			}
			wg.Done() // hey waitgroup, this goroutine is done
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
