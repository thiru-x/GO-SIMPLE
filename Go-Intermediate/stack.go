package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1 // Stack empty
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}
func main() {
	var stack Stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	fmt.Println("Top element:", stack.Peek()) // Should print 40
	fmt.Println("Poped element:", stack.Pop())
	fmt.Println("Top element:", stack.Peek())           // Should print 30
	stack.Pop()                                         // Pop the top element
	fmt.Println("Top element after pop:", stack.Peek()) // Should print 20
	stack.Pop()                                         // Pop the top element
	fmt.Println("Top element after pop:", stack.Peek()) // Should print 10
	stack.Pop()                                         // Pop the last element
	fmt.Println("Top element after pop:", stack.Peek()) // Should print -1 (stack empty)
}
