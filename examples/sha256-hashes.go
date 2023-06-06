// sha256-hashes.go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"
	h := sha256.New()

	// convert string to bytes
	h.Write([]byte(s))

	// gets finalized hash result as a byte slice
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
