// waitgroups.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// I would never have thought of this
		i := i

		go func() {
			// Done is the opposite of Add(1)
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
