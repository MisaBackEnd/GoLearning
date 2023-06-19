package main

import "errors"

// Datastructure in which you can only add and delete items from the array from one end.
// This end is called top. last in / first out (lifo)
type Stack struct {
	stack []int
}

// push: add element at the end of the array
func (s *Stack) Push(element int) {
	s.stack = append(s.stack, element)
}

// pop: remove the last added element of the array
func (s *Stack) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("The stack doesn't have elements")
	}
	data := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return data, nil
}

func main() {}
