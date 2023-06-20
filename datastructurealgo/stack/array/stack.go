package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	index := len(s.data) - 1
	value := s.data[index]
	s.data = s.data[:index]

	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	index := len(s.data) - 1
	value := s.data[index]

	return value, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}

func main() {
	stack := NewStack()

	// Push elements onto the stack
	stack.Push(5)
	stack.Push(10)
	stack.Push(3)

	// Pop elements from the stack
	value, _ := stack.Pop()
	fmt.Println("Popped Value:", value) // Output: 3

	// Peek at the top element of the stack
	value, _ = stack.Peek()
	fmt.Println("Top Value:", value) // Output: 10

	// Check if the stack is empty
	fmt.Println("Is Empty:", stack.IsEmpty()) // Output: false

	// Get the size of the stack
	fmt.Println("Size:", stack.Size()) // Output: 2
}
