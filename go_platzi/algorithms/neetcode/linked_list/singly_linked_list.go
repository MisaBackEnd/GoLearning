package main

import "fmt"

// list node data structure
type ListNode struct {
	Value int
	Next  *ListNode
}

// list node constructor
func NewListNode(val int) *ListNode {
	return &ListNode{
		Value: val,
		Next:  nil,
	}
}

// singly linked list data structure
type SinglyLinkedList struct {
	Head, Tail *ListNode
}

// new singly linked list constructor
func NewSinglyLinkedList() *SinglyLinkedList {
	node := NewListNode(0)
	return &SinglyLinkedList{
		Head: node,
		Tail: node,
	}
}

// insert value at the end of a singly linked list
func (s *SinglyLinkedList) InsertEnd(val int) {
	s.Tail.Next = NewListNode(val)
	s.Tail = s.Tail.Next
}

// remove value at index of a singly linked list
func (s *SinglyLinkedList) Remove(index int) {
	c := 1
	node := s.Head
	// the node is the second element
	for c < index {
		node = node.Next
		c++
	}
	fmt.Println(node)
	// the index == 0 make the next node the head
	if node != nil {
		switch {
		case s.Head == node && node.Next != nil && index == 0:
			s.Head = node.Next
		case node.Next == s.Tail:
			s.Tail = node
		default:
			if node.Next != nil {
				node.Next = node.Next.Next
			}
		}
	}
}

func (s *SinglyLinkedList) Print() {
	node := s.Head
	for node != nil {
		fmt.Printf(" %d ->", node.Value)
		node = node.Next
	}
}

func main() {
	sl := NewSinglyLinkedList()
	sl.InsertEnd(5)
	sl.InsertEnd(6)
	sl.Remove(0)
	sl.Print()
}
