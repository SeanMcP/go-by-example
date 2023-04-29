package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("idx:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	// note: order is not guaranteed
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// this block is not in the example
	for _, v := range kvs {
		fmt.Println("val:", v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
