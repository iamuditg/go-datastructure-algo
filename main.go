package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (ll *LinkedList) Insert(value int) {
	node := &Node{
		value: value,
	}

	current := ll.head
	if ll.head == nil {
		ll.head = node
	} else {
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	ll.size++
}

func (ll *LinkedList) InsertElementAt(value int, index int) {
	if index < 0 {
		fmt.Println("index is invalid")
		return
	}
	node := &Node{value: value}
	current := ll.head
	if index == 0 {
		node.next = current
		ll.head = node
	} else {
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		node.next = current.next
		current.next = node
	}
	ll.size++
}

func (ll *LinkedList) printLinkedList() {
	if ll.head == nil {
		fmt.Println("linked list is empty")
		return
	}
	current := ll.head
	for current != nil {
		fmt.Println(current.value, " ")
		current = current.next
	}
}

func (ll *LinkedList) deleteElementAt(index int) {
	if index < 0 {
		fmt.Println("index is invalid")
		return
	}
	current := ll.head
	if index == 0 {
		current = current.next
	} else {
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}
	ll.size--
}

func (ll *LinkedList) ReverseLinkedList() {
	current := ll.head
	var prev *Node = nil
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	ll.head = prev
}

func (ll *LinkedList) SearchElement(value int) bool {
	if ll.head == nil {
		fmt.Println("no element in linked list")
		return false
	}
	current := ll.head
	for current.next != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func (ll *LinkedList) MiddleElement() *Node {
	slow := ll.head
	fast := ll.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func MergeSortedLinkedList(list1, list2 *Node) *Node {
	dummy := &Node{}
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

func main() {
	ll := &LinkedList{}
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(6)
	ll.Insert(9)

	// print
	ll.printLinkedList()

	// insertAt
	ll.InsertElementAt(5, 2)
	println()
	//print
	ll.printLinkedList()

	println()
	// deleteAt
	ll.deleteElementAt(2)
	//print
	ll.printLinkedList()

	//println()
	////reverse
	//ll.ReverseLinkedList()
	//ll.printLinkedList()

	//println()
	//fmt.Println(ll.SearchElement(13))

	fmt.Println(ll.MiddleElement())

	ll2 := &LinkedList{}
	ll2.Insert(1)
	ll2.Insert(4)
	ll2.Insert(6)
	ll2.Insert(9)

	list := MergeSortedLinkedList(ll.head, ll2.head)
	for list != nil {
		fmt.Println(list.value)
		list = list.next
	}
}
