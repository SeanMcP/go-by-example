// non-blocking-channel-operations.go
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("recieved message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	// no buffer, no receiver, so no send
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	// no send, so no receive
	case msg := <-messages:
		fmt.Println("received message", msg)
	// no send, so no receive
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
