package main

import "fmt"

func main() {
	// make a map with string keys and int values
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	// you don't need to create a variable to print,
	// but that's what the example did.
	fmt.Println("v1: ", v1)

	// returns "zero value" if key doesn't exist
	// see https://go.dev/ref/spec#The_zero_value
	v3 := m["k3"]
	fmt.Println("v3: ", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// prs is a "does key exist" boolean
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
