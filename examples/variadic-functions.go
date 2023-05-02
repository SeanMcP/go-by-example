// variadic functions can be called with any number of trailing arguments

package main

import "fmt"

func sum(nums ...int) {
	// inline print
	// fmt.Print(nums)
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(nums, total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// note: apply ... _after_ slice
	sum(nums...)
}
