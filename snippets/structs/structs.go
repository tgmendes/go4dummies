package main

import "fmt"

// START OMIT
type person struct {
	name string
	age  int
}

// A method of person struct, accepting no arguments
// and returning an int.
func (p person) birthYear() int { // HL
	return 2019 - p.age
}

func main() {
	p := person{
		name: "Alice",
		age:  31,
	}
	fmt.Println(person{"Bob", 20})
	fmt.Println(p)
	fmt.Println(p.birthYear())
}

// END OMIT
