// command-line-flags.go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag name, default value, description
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// use existing variable
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

/**

examples/command-line-flags -word=opt -numb=7 -fork -svar=flag
examples/command-line-flags -word=opt
examples/command-line-flags -word=opt a1 a2 a3
examples/command-line-flags -word=opt a1 a2 a3 -numb=7
examples/command-line-flags -h
examples/command-line-flags -wat

*/
