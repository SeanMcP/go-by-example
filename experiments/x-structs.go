package main

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Increment(value ...int) {
	if value != nil {
		c.count += value[0]
	} else {
		c.count++
	}
}

func main() {
	c1 := Counter{}
	c1.Increment()
	fmt.Println("c1:", c1)
	c1.Increment(5)
	fmt.Println("c1:", c1)
}
