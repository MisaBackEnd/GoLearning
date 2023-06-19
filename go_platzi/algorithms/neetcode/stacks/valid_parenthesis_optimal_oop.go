package main

type Stack struct {
	data []byte
}

func NewStack() *Stack {
	return &Stack{
		data: []byte{},
	}
}

func (s *Stack) Push(value byte) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(s.data) - 1
	value := s.data[index]
	s.data = s.data[:index]
	return value, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	var letter byte
	closedOpen := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := NewStack()
	for i := 0; i < len(s); i++ {
		letter = s[i]

		switch letter {
		case '(', '[', '{':
			stack.Push(letter)

		case ')', ']', '}':
			top, ok := stack.Pop()
			if !ok {
				return false
			}
			pair := closedOpen[letter]
			if pair != top {
				return false
			}
		}
	}
	if stack.IsEmpty() {
		return true
	}
	return false
}
