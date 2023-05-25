// regular-expressions.go
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach") // boolean
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	// first match
	fmt.Println(r.FindString("peach punch")) // string

	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch")) // [match, group]

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// all matches
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// "find all. jk, just 2"
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
