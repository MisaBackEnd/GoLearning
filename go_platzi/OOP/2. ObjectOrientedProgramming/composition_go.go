package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id     int
	salary float32
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (fe FullTimeEmployee) getMessage() string {
	return "This is a full time employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (te TemporaryEmployee) getMessage() string {
	return "This is a temporary employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(obj PrintInfo) {
	fmt.Println(obj.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Misael"
	ftEmployee.age = 34
	ftEmployee.id = 1
	ftEmployee.salary = 4500000

	tEmployee := TemporaryEmployee{}

	getMessage(tEmployee)
	getMessage(ftEmployee)
}
