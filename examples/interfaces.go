package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// as of right now, rect and circle have no relationship to geometry

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// rect now implements geometry

func (c circle) area() float64 {
	// how to do powers with math?
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// circle now implements geometry

func measure(g geometry) {
	fmt.Println(g)
	// logging differs slightly from example
	fmt.Println("  area: ", g.area())
	fmt.Println("  perim:", g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
