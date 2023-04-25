package main

import (
	"fmt"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{value: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	last := l.head
	for last.next != nil {
		last = last.next
	}
	last.next = newNode
}

func (l *LinkedList) Print() {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	// Collect the digits in a slice
	digits := []string{}
	ptr := l.head
	for ptr != nil {
		digits = append(digits, fmt.Sprintf("%d", ptr.value))
		ptr = ptr.next
	}

	// Print the digits in reverse order
	fmt.Println(strings.Join(digits, ""))
}

func addReverse(n1, n2 *Node) *Node {
	head := &Node{}
	current := head
	carry := 0

	for n1 != nil || n2 != nil || carry > 0 {
		sum := carry

		if n1 != nil {
			sum += n1.value
			n1 = n1.next
		}

		if n2 != nil {
			sum += n2.value
			n2 = n2.next
		}

		carry = sum / 10
		sum = sum % 10

		newNode := &Node{value: sum}
		current.next = newNode
		current = newNode
	}

	return head.next
}

func main() {
	l1 := &LinkedList{}
	l1.Append(7)
	l1.Append(1)
	l1.Append(6)

	l2 := &LinkedList{}
	l2.Append(5)
	l2.Append(9)
	l2.Append(2)

	sum := addReverse(l1.head, l2.head)

	// Print the sum
	l1.Print()
	fmt.Print(" + ")
	l2.Print()
	fmt.Print(" = ")
	sumList := &LinkedList{head: sum}
	sumList.Print()
}
