// range-over-channels.go
package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// no more sends, but receives okay
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
