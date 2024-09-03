package main

import "fmt"

type person struct {
	name string
	age  int
}

// p is a copy of the pointer to a person struct
func update(p *person) {
	// Dereference the pointer and point to a new address in memory where where a new person struct with name "New Name" and age 30 is created
	// However, this does not modify the original person struct that was passed as an argument to the function.
	// The original person struct is still at the same address in memory.
	p = &person{"New Name", 30}
}

func main() {
	p1 := person{"John", 25}
	p2 := p1
	update(&p2)
	fmt.Println(p1) // Output: {John 25}
	fmt.Println(p2) // Output: {New Name 30}

	// Pointers
	var a = 5
	var p = &a // p holds variable a's memory address
	fmt.Printf("Address of var a: %p\n", p)
	fmt.Printf("Value of var a: %v\n", *p)

	// Let's change a value (using the initial variable or the pointer)
	*p = 3 // using pointer
	a = 3  // using initial var

	fmt.Printf("Address of var a: %p\n", p)
	fmt.Printf("Value of var a: %v\n", *p)
}
