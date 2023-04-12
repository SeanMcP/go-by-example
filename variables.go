package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// Zero-valued
	var e int
	fmt.Println(e)

	// := syntax for declaring & initializing
	f := "apple"
	fmt.Println(f)
}
