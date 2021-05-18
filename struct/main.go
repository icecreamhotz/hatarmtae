package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	p := person{
		"Anucha",
		"P",
	}
	fmt.Printf("%v", p.print())
}

func (p person) print() person {
	return p
}
