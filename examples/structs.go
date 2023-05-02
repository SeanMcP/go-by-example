package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	// why return a pointer?
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Anne", age: 40})
	// idiomatic to use constructor func
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 31} // different from example
	fmt.Println(s.name)

	// interesting that it doesn't seem to matter if you're dealing with a reference or a pointer
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
