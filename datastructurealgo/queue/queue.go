package main

import "fmt"

type queue []int

func (q *queue) Empty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

func (q *queue) enqueue(element int) {
	*q = append(*q, element)
}

func (q *queue) dequeue() (int, bool) {
	if q.Empty() {
		return 0, true
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, false
	}
}

func main() {
	var queue queue
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(5)
	for len(queue) > 0 {
		dequeue, b := queue.dequeue()
		if b == false {
			fmt.Println(dequeue)
		}
	}
}
