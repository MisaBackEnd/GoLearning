package main

import "fmt"

type animal interface {
	move() string
}

type dog struct{}
type fish struct{}
type bird struct{}

func (dog) move() string {
	return "I'm a dog that is walking"
}
func (fish) move() string {
	return "I'm a fish that is swimming"
}
func (bird) move() string {
	return "I'm a bird that is flying"
}

func moveAnimal(an animal) {
	fmt.Println(an.move())
}

func main() {
	doggy := dog{}
	moveAnimal(doggy)
	fishy := fish{}
	moveAnimal(fishy)
	birdy := bird{}
	moveAnimal(birdy)
}
