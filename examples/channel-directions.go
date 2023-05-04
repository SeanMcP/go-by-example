package main

import "fmt"

// send only, error on receive
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings: receive only
// pongs: send only
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// sends msg to pings
	ping(pings, "passed message")
	// receives msg from pings, sends to pongs
	pong(pings, pongs)

	// receive msg from pongs
	fmt.Println(<-pongs)
}
