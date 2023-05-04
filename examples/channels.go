package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
		// different from example
		messages <- "pong"
	}()

	msg := <-messages
	fmt.Println(msg)

	// different from example
	msg2 := <-messages
	fmt.Println(msg2)
}
