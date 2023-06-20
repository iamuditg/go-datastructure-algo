package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top  *Node
	size int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value int) {
	newNode := &Node{value: value, next: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	value := s.top.value
	s.top = s.top.next
	s.size--

	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.top.value, nil
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
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
