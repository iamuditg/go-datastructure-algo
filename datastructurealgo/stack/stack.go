package main

import (
	"errors"
	"fmt"
	"log"
)

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{data: make([]interface{}, 0)}
}

func (s *Stack) push(item interface{}) {
	s.data = append(s.data, item)
}

func (s *Stack) pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, errors.New("stack is empty")
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item, nil
}

func (s *Stack) peek() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func main() {
	stack := NewStack()
	stack.push(10)
	stack.push(20)
	stack.push(30)
	stack.push(40)

	pop, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pop)

	peek, err := stack.peek()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(peek)
}
