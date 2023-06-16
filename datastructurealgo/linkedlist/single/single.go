package main

import "fmt"

// ListNode represents a node in the linked list
type ListNode struct {
	value int
	next  *ListNode
}

// LinkedList represents a singly linked list
type LinkedList struct {
	head *ListNode
	size int
}

// InsertElement inserts a new elements at the end of the linked lists
func (ll *LinkedList) InsertElement(value int) {
	newNode := &ListNode{value: value}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.size++
}

// InsertElementAt inserts a new element at the specified index position on the Linked List
func (ll *LinkedList) InsertElementAt(value, index int) {
	if index == 0 || index > ll.size {
		fmt.Println("invalid index position")
		return
	}
	newNode := &ListNode{value: value}

	if index == 0 {
		newNode.next = ll.head
		ll.head = newNode
	} else {
		current := ll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
	}
	ll.size++
}

// PrintLinkedList prints the elements of the linked list
func (ll *LinkedList) PrintLinkedList() {
	current := ll.head
	for current != nil {
		fmt.Println(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := &LinkedList{}
	// Insert element in Linked list
	list.InsertElement(3)
	list.InsertElement(4)
	list.InsertElement(5)
	list.InsertElement(6)
	list.InsertElement(7)

	// InsertAt
	list.InsertElementAt(8, 2)

	// Print Linked List
	list.PrintLinkedList()

}
