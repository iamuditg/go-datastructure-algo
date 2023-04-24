package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) insert(value int) {
	newNode := Node{}
	newNode.data = value
	if l.head == nil {
		l.head = &newNode
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &newNode
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) printList() {
	if l.len == 0 || l.len < 0 {
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println(ptr.data)
		ptr = ptr.next
	}
}

func (l *LinkedList) insertAt(pos int, val int) {
	if l.len < 0 {
		return
	}
	newNode := Node{}
	newNode.data = val
	if l.len == 0 && l.head == nil {
		l.head = &newNode
		l.len++
		return
	}
	nextNode := l.GetAt(pos)
	if nextNode == nil {
		return
	}
	newNode.next = nextNode
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}

func (l *LinkedList) GetAt(pos int) *Node {
	if pos < 0 || l.len <= 0 {
		return nil
	}
	ptr := l.head
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (l *LinkedList) Search(val int) {
	if l.len <= 0 {
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.data == val {
			fmt.Printf("search value is found %d in pos %d \n", val, i)
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) DeleteAt(pos int) {
	if l.len <= 0 || pos <= 0 {
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if i == pos {
			prevNode := l.GetAt(pos - 1)
			prevNode.next = ptr.next
			l.len--
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) DeleteDuplicate() {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		runner := current
		for runner != nil && runner.next != nil {
			if runner.next.data == current.data {
				runner.next = runner.next.next
				l.len--
			} else {
				runner = runner.next
			}
		}
		current = current.next
	}
}

func (l *LinkedList) findKthToLast(k int) *Node {
	p1, p2 := l.head, l.head

	for i := 0; i < k-1; i++ {
		if p2 == nil {
			return nil
		}
		p2 = p2.next
	}

	for p2.next != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p1
}

func (l *LinkedList) DeleteMidElement() {
	if l.head == nil || l.len == 0 || l.head.next == nil {
		return
	}
	slow, fast := l.head, l.head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	prev := l.head
	for prev.next != slow {
		prev = prev.next
	}
	prev.next = slow.next
	l.len--
}

func main() {
	l := LinkedList{}
	l.insert(10)
	l.insert(20)
	l.insert(20)
	l.insert(30)
	l.insert(30)
	l.insert(40)
	l.DeleteDuplicate()
	last := l.findKthToLast(4)
	fmt.Println(last)
	l.DeleteMidElement()
	l.printList()

}
