package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	front *Node
	rear  *Node
	size  int
}

func NewQueue() *Queue {
	return &Queue{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

func (q *Queue) Enqueue(value int) {
	newNode := &Node{value: value, next: nil}
	if q.isEmpty() {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size++
}

func (q *Queue) Dequeue() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("Queue is empty")
	}
	value := q.front.value
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	q.size--
	return value, nil
}

func (q *Queue) Front() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("Queue is empty")
	}
	return q.front.value, nil
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func (q *Queue) Print() {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return
	}
	current := q.front
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
	fmt.Println()
}

func main() {
	queue := NewQueue()

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)

	queue.Print() // Output: 10 20 30 40

	value, err := queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dequeued value:", value) // Output: Dequeued value: 10
	}

	frontValue, err := queue.Front()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Front value:", frontValue) // Output: Front value: 20
	}
}
