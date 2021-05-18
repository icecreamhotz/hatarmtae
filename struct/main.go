package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contract  struct {
		email   string
		zipcode string
	}
}

func main() {
	p := person{
		"Anucha",
		"P",
		struct {
			email   string
			zipcode string
		}{
			email:   "admin@hotmail.com",
			zipcode: "001",
		},
	}
	fmt.Printf("%v", p.print())
}

func (p person) print() person {
	return p
}
