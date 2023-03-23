package main

import "fmt"

type stack []string

// IsEmpty Check if Stack is Empty
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Pop Remove and return top element of stack. Return false if stack is empty
func (s *stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element
		element := (*s)[index] //Index into the slices and obtain the element
		*s = (*s)[:index]      // Remove it from the stack by slicing it off
		return element, true
	}
}

func main() {
	var stack stack
	stack.Push("this")
	stack.Push("is")
	stack.Push("sparta!")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}
