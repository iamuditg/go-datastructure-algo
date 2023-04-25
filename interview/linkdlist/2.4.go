package linkdlist

func (l *LinkedList) Partition(x int) {
	if l == nil || l.head == nil {
		return
	}

	var smaller, larger *Node
	var smallerTail, largerTail *Node

	// Partition the list
	cur := l.head
	for cur != nil {
		if cur.data < x {
			if smaller == nil {
				smaller = cur
				smallerTail = smaller
			} else {
				smallerTail.next = cur
				smallerTail = cur
			}
		} else {
			if larger == nil {
				larger = cur
				largerTail = larger
			} else {
				largerTail.next = cur
				largerTail = cur
			}
		}
		cur = cur.next
	}

	// Merge the smaller and larger lists
	if smaller == nil {
		l.head = larger
	} else {
		l.head = smaller
		smallerTail.next = larger
	}
	largerTail.next = nil

}
