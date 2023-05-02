package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

// this block is not in the example
func moreVals() (int, int, int) {
	return 1, 2, 3
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// this block is not in the example
	x, y, z := moreVals()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	_, c := vals()
	fmt.Println(c)
}
