package main

import "fmt"

func main() {
	// messages := make(chan string)
	// fatal error: all goroutines are asleep - deadlock!
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"
	// messages <- "test"

	// q: why is it different if the sends come from a goroutine?

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
