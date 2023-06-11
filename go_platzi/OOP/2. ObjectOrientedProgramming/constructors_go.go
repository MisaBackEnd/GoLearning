package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func NewEmployee(id int, name string) *Employee {
	return &Employee{
		id:   id,
		name: name,
	}
}

func main() {
	// without constructor, or implicit
	e := Employee{
		id:   1,
		name: "Misael",
	}
	fmt.Printf("%+v", e)

	// using new, returns a reference of the e2 employee
	e2 := new(Employee)
	fmt.Printf("%+v", e2)
	fmt.Printf("%+v", *e2)

	e3 := NewEmployee(2, "Katherine")
	fmt.Printf("%+v", e3)
	fmt.Printf("%+v", *e3)

}
