// panic.go
package main

import "os"

func main() {
	// panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		// q. why didn't this fire? why would it?
		panic(err)
	}
}
