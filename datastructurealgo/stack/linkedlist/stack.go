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
	start *Node
	top   *Node
	size  int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value int) {
	newNode := &Node{value: value}

	if s.start == nil {
		s.start = newNode
		s.top = newNode
	} else {
		s.top.next = newNode
		s.top = newNode
	}

	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.top == nil {
		return 0, errors.New("Stack is empty")
	}

	value := s.top.value

	if s.start == s.top {
		s.start = nil
		s.top = nil
	} else {
		curr := s.start
		for curr.next != s.top {
			curr = curr.next
		}
		curr.next = nil
		s.top = curr
	}

	s.size--
	return value, nil
}

func (s *Stack) Print() {
	curr := s.start
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func (s *Stack) Peek() (int, error) {
	if s.top == nil {
		return 0, errors.New("Stack is empty")
	}

	return s.top.value, nil
}

func (s *Stack) MidElement() (int, error) {
	if s.size == 0 {
		return 0, errors.New("Stack is empty")
	}

	mid := s.size / 2
	curr := s.start

	for i := 0; i < mid; i++ {
		curr = curr.next
	}

	return curr.value, nil
}

func main() {
	stack := NewStack()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	stack.Print() // Output: 10 20 30 40

	value, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Popped value:", value) // Output: Popped value: 40
	}

	value, err = stack.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Top value:", value) // Output: Top value: 30
	}

	midValue, err := stack.MidElement()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Middle value:", midValue) // Output: Middle value: 20
	}
}
