// sorting-by-functions.go
package main

import (
	"fmt"
	"sort"
)

type byLength []string

// names the methods required by sort.Interface

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	// q: calling a type?
	sort.Sort(byLength(fruits))
	// sorts in place
	fmt.Println(fruits)
}
