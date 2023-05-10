// worker-pools.go
package main

import (
	"fmt"
	"time"
)

func worker(id string, jobs <-chan int, results chan<- int) {
	// q: we aren't receiving from jobs, so why aren't jobs repeated?
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	ids := []string{"a", "b", "c"}
	for _, id := range ids {
		go worker(id, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
