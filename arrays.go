package main

import "fmt"

func main() {
	var a [5]int
	// Why "emp"?
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// len is a builtin
	fmt.Println("len:", len(a))

	// Interesting syntax here
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array of int arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
