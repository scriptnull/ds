package adt

import "errors"

// Queue - rear inserts and front removals

// ArrQueue is queue implementation with fixed size array
type ArrQueue struct {
	arr         [5]int
	front       int
	rear        int
	currentSize int
}

// Enqueue inserts elements to the queue in rear position
func (q *ArrQueue) Enqueue(element int) error {
	if q.currentSize == len(q.arr) {
		return errors.New("Queue Overflow")
	}

	// insert element at rear position
	q.arr[q.rear] = element

	// increment rear position circularly
	q.rear++
	q.rear = q.rear % len(q.arr)

	q.currentSize++

	return nil
}

// Dequeue removes elements from the queue in the front position
func (q *ArrQueue) Dequeue() (int, error) {
	if q.currentSize == 0 {
		return 0, errors.New("Queue Underflow")
	}

	// fetch the element at front
	element := q.arr[q.front]

	// increment front position circularly
	q.front++
	q.front = q.front % len(q.arr)

	q.currentSize--

	return element, nil
}

// Peek returns element at the front
func (q *ArrQueue) Peek() (int, error) {
	if q.currentSize == 0 {
		return 0, errors.New("Cannot Peak on empty queue")
	}
	return q.arr[q.front], nil
}

// DynQueue is queue implementation with dynamic array
type DynQueue struct {
	arr []int
}

// Enqueue inserts elements to the queue in rear position
func (q *DynQueue) Enqueue(element int) error {
	q.arr = append(q.arr, element)
	return nil
}

// Dequeue removes elements from the queue in the front position
func (q *DynQueue) Dequeue() (int, error) {
	if len(q.arr) == 0 {
		return 0, errors.New("Queue underflow")
	}

	front := q.arr[0]
	q.arr = q.arr[1:]

	return front, nil
}

// Peek returns element at the front
func (q *DynQueue) Peek() (int, error) {
	if len(q.arr) == 0 {
		return 0, errors.New("Cannot Peak on empty queue")
	}
	return q.arr[0], nil
}
