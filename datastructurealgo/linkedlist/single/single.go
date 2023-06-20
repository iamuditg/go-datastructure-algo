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

// DeleteElementAt deletes the element at the specified index position in the linked list
func (ll *LinkedList) DeleteElementAt(index int) {
	if index < 0 || index >= ll.size {
		fmt.Println("Invalid index position")
		return
	}

	if index == 0 {
		ll.head = ll.head.next
	} else {
		current := ll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}
	ll.size--
}

// DeleteAllElements deletes all elements in the linked list
func (ll *LinkedList) DeleteAllElements() {
	ll.head = nil
	ll.size = 0
}

// ReverseLinkedList reverse the linked list
func (ll *LinkedList) ReverseLinkedList() {
	var prev *ListNode = nil
	current := ll.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	ll.head = prev
}

// MergeSortedLinkedList merges two sorted linked lists into a single sorted linked list
func MergeSortedLinkedList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.value <= list2.value {
			current.next = list1
			list1 = list1.next
		} else {
			current.next = list2
			list2 = list2.next
		}
		current = current.next
	}

	if list1 != nil {
		current.next = list1
	} else {
		current.next = list2
	}

	return dummy.next
}

// SearchElement searches for a given element in the linked list and returns true if found, false otherwise
func (ll *LinkedList) SearchElement(value int) bool {
	current := ll.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

// MiddleElement returns the middle element of the linked list
func (ll *LinkedList) MiddleElement() *ListNode {
	slow := ll.head
	fast := ll.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
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
	// Create a new linked list
	ll := LinkedList{}

	// Insert elements
	ll.InsertElement(1)
	ll.InsertElement(2)
	ll.InsertElement(3)
	ll.InsertElement(4)
	ll.InsertElement(5)

	// Print the linked list
	fmt.Println("Linked List:")
	ll.PrintLinkedList() // Output: 1 2 3 4 5

	// Insert element at index
	ll.InsertElementAt(10, 2)

	// Print the modified linked list
	fmt.Println("Modified Linked List:")
	ll.PrintLinkedList() // Output: 1 2 10 3 4 5

	// Delete element at index
	ll.DeleteElementAt(3)

	// Print the modified linked list
	fmt.Println("Modified Linked List:")
	ll.PrintLinkedList() // Output: 1 2 10 4 5

	// Reverse the linked list
	ll.ReverseLinkedList()

	// Print the reversed linked list
	fmt.Println("Reversed Linked List:")
	ll.PrintLinkedList() // Output: 5 4 10 2 1

	// Create another linked list
	ll2 := LinkedList{}

	// Insert elements
	ll2.InsertElement(3)
	ll2.InsertElement(6)
	ll2.InsertElement(9)

	// Print the second linked list
	fmt.Println("Second Linked List:")
	ll2.PrintLinkedList() // Output: 3 6 9

	// Merge the two linked lists in sorted order
	merged := MergeSortedLinkedList(ll.head, ll2.head)

	// Print the merged linked list
	fmt.Println("Merged Linked List:")
	for merged != nil {
		fmt.Print(merged.value, " ")
		merged = merged.next
	}
	fmt.Println() // Output: 1 2 3 4 5 6 9

	// Search for an element in the linked list
	fmt.Println("Search Result:", ll.SearchElement(4)) // Output: true
	fmt.Println("Search Result:", ll.SearchElement(7)) // Output: false

	// Get the middle element of the linked list
	middle := ll.MiddleElement()
	fmt.Println("Middle Element:", middle.value) // Output: 10
}
