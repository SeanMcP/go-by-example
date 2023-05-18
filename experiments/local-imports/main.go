package main

import (
	"fmt"

	"main/sub"
)

func main() {
	fmt.Println("Hello from outer!")
	sub.Hello()
}
