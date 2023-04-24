package linkdlist

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
