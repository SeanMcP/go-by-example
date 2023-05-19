// recover.go
package main

import "fmt"

func mayPanic() {
	// jk always panics
	panic("a problem")
}

func main() {
	// recover only works in deferred functions
	defer func() {
		// q. why r and not e/err?
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// will not run. main execution stops after panic
	fmt.Println("After mayPanic()")
}
