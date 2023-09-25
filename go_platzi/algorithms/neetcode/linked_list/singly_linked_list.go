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
func (s *SinglyLinkedList) Find(index int) (*ListNode, *ListNode) {
	node := s.Head
	var prev *ListNode
	for i := 0; i < index; i++ {
		prev = node
		node = node.Next
	}
	return prev, node
}

func (s *SinglyLinkedList) Remove(index int) {
	prev, node := s.Find(index)
	if index == 0 {
		s.Head = node.Next
	} else {
		prev.Next = node.Next
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
	sl.Remove(3)
	sl.Print()
}
