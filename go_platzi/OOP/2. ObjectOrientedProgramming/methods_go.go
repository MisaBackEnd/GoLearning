package main

import "fmt"

type Employee struct {
	id   int
	name string
}

// receiver function where the
// with *Employee we're saying that
// this is a method of a concrete object
// e where its type is Employee.
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	e.id = 0
	e.name = "Tomas"
	fmt.Printf("%+v", e)
	e.SetId(1)
	fmt.Printf("%+v\n", e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

}
