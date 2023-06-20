package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	data     []int
	capacity int
	front    int
	rear     int
	size     int
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		data:     make([]int, capacity),
		capacity: capacity,
		front:    0,
		rear:     -1,
		size:     0,
	}
}

func (q *Queue) Enqueue(value int) error {
	if q.isFull() {
		return errors.New("Queue is full")
	}
	q.rear = (q.rear + 1) % q.capacity
	q.data[q.rear] = value
	q.size++
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("Queue is empty")
	}
	value := q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return value, nil
}

func (q *Queue) Front() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("Queue is empty")
	}
	return q.data[q.front], nil
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func (q *Queue) isFull() bool {
	return q.size == q.capacity
}

func (q *Queue) Print() {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return
	}
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		fmt.Printf("%d ", q.data[index])
	}
	fmt.Println()
}

func main() {
	queue := NewQueue(5)

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
