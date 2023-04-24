package linkdlist

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
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
