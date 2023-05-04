package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	// q: do we need the 1?
	// a: no, but maybe it's best practice
	done := make(chan bool, 1)

	go worker(done)

	// block until we receive a send from the goroutine
	// otherwise the program ends before the worker runs
	<-done
}
