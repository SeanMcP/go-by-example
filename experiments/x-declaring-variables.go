package main

import "fmt"

func main() {
	var v string = "a"
	v += "b"
	fmt.Println("var", v)

	var v2 = "a"
	v2 += "b"
	fmt.Println("v2 ", v2)

	var v3 string
	v3 += "!"

	s := "a"
	s += "b"
	fmt.Println("svd", s)

	const c = "a"
	fmt.Println("const", c)
}
