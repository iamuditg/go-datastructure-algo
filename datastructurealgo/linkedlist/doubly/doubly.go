package main

import (
	"fmt"
)

// Node represents a node in the doubly linked list.
type Node struct {
	value int
	prev  *Node
	next  *Node
}

// DoublyLinkedList represents a doubly linked list.
type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

// NewDoublyLinkedList creates a new instance of DoublyLinkedList.
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// InsertElement inserts an element at the end of the linked list.
func (dll *DoublyLinkedList) InsertElement(value int) {
	newNode := &Node{value: value}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}

	dll.size++
}

// InsertAt inserts an element at the specified index in the linked list.
func (dll *DoublyLinkedList) InsertAt(value int, index int) {
	if index < 0 || index > dll.size {
		return
	}

	newNode := &Node{value: value}

	if index == 0 {
		if dll.head == nil {
			dll.head = newNode
			dll.tail = newNode
		} else {
			newNode.next = dll.head
			dll.head.prev = newNode
			dll.head = newNode
		}
	} else if index == dll.size {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	} else {
		curr := dll.head
		for i := 0; i < index-1; i++ {
			curr = curr.next
		}

		newNode.next = curr.next
		newNode.prev = curr
		curr.next.prev = newNode
		curr.next = newNode
	}

	dll.size++
}

// DeleteAt deletes the element at the specified index from the linked list.
func (dll *DoublyLinkedList) DeleteAt(index int) {
	if index < 0 || index >= dll.size {
		return
	}

	if index == 0 {
		if dll.size == 1 {
			dll.head = nil
			dll.tail = nil
		} else {
			dll.head = dll.head.next
			dll.head.prev = nil
		}
	} else if index == dll.size-1 {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	} else {
		curr := dll.head
		for i := 0; i < index; i++ {
			curr = curr.next
		}

		curr.prev.next = curr.next
		curr.next.prev = curr.prev
	}

	dll.size--
}

// DeleteAll deletes all elements from the linked list.
func (dll *DoublyLinkedList) DeleteAll() {
	dll.head = nil
	dll.tail = nil
	dll.size = 0
}

// ReverseLinkedList reverses the order of elements in the linked list.
func (dll *DoublyLinkedList) ReverseLinkedList() {
	curr := dll.head
	var prev *Node

	for curr != nil {
		next := curr.next
		curr.next = prev
		curr.prev = next
		prev = curr
		curr = next
	}

	dll.head = prev
}

// SearchElement searches for an element in the linked list and returns true if found, false otherwise.
func (dll *DoublyLinkedList) SearchElement(value int) bool {
	curr := dll.head

	for curr != nil {
		if curr.value == value {
			return true
		}
		curr = curr.next
	}

	return false
}

// MiddleElement returns the middle element of the linked list.
func (dll *DoublyLinkedList) MiddleElement() *Node {
	if dll.head == nil {
		return nil
	}

	slow := dll.head
	fast := dll.head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	return slow
}

func main() {
	dll := NewDoublyLinkedList()

	// Insert elements
	dll.InsertElement(5)
	dll.InsertElement(10)
	dll.InsertElement(3)

	// Insert an element at a specific index
	dll.InsertAt(7, 1)

	// Delete an element at a specific index
	dll.DeleteAt(2)

	// Reverse the doubly linked list
	dll.ReverseLinkedList()

	// Search for an element in the doubly linked list
	fmt.Println("Search Result:", dll.SearchElement(4)) // Output: true
	fmt.Println("Search Result:", dll.SearchElement(7)) // Output: false

	// Get the middle element of the doubly linked list
	middle := dll.MiddleElement()
	fmt.Println("Middle Element:", middle.value) // Output: 10
}
