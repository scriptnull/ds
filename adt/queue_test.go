package adt

import "testing"

func TestArrayImplementation(t *testing.T) {
	var queue = &ArrQueue{}

	_, err := queue.Dequeue()
	if err == nil {
		t.Error("Expects Queue underflow error")
	}

	_, err = queue.Peek()
	if err == nil {
		t.Error("Expects error on reading empty queue")
	}

	err = queue.Enqueue(1)
	if err != nil {
		t.Error("Expects to insert element in queue, but got error ", err)
	}
	err = queue.Enqueue(2)
	err = queue.Enqueue(3)
	err = queue.Enqueue(4)
	err = queue.Enqueue(5)
	err = queue.Enqueue(6)
	if err == nil {
		t.Error("Expects queue overflow error")
	}

	// Dequeue two elements
	front, err := queue.Dequeue()
	if err != nil {
		t.Error("Expected dequeue to happen without error, but got", err)
	}
	if front != 1 {
		t.Error("Expected 1, got ", front)
	}
	front, err = queue.Dequeue()
	if err != nil {
		t.Error("Expected dequeue to happen without error, but got", err)
	}
	if front != 2 {
		t.Error("Expected 2, got ", front)
	}

	// Enqueue Two Elements
	err = queue.Enqueue(6)
	if err != nil {
		t.Error("Expects to insert element in queue, but got error ", err)
	}
	err = queue.Enqueue(7)
	if err != nil {
		t.Error("Expects to insert element in queue, but got error ", err)
	}

	front, err = queue.Peek()
	if err != nil {
		t.Error("Expects peek to happen without error, but got", err)
	}
	if front != 3 {
		t.Error("Expected 3 to be at front, but got", front)
	}

	// Dequeue 4 Elements
	front, err = queue.Dequeue()
	if err != nil {
		t.Error("Expected dequeue to happen without error, but got", err)
	}
	if front != 3 {
		t.Error("Expected 1, got ", front)
	}
	front, err = queue.Dequeue()
	if err != nil {
		t.Error("Expected dequeue to happen without error, but got", err)
	}
	if front != 4 {
		t.Error("Expected 4, got ", front)
	}
	front, err = queue.Dequeue()
	if err != nil {
		t.Error("Expected dequeue to happen without error, but got", err)
	}
	if front != 5 {
		t.Error("Expected 5, got ", front)
	}
	front, err = queue.Dequeue()
	if err != nil {
		t.Error("Expected dequeue to happen without error, but got", err)
	}
	if front != 6 {
		t.Error("Expected 6, got ", front)
	}

	front, err = queue.Peek()
	if err != nil {
		t.Error("Expects peek to happen without error, but got", err)
	}
	if front != 7 {
		t.Error("Expected 3 to be at front, but got", front)
	}

	// Empty the queue
	front, err = queue.Dequeue()
	if err != nil {
		t.Error("Expected dequeue to happen without error, but got", err)
	}
	if front != 7 {
		t.Error("Expected 7, got ", front)
	}

	_, err = queue.Dequeue()
	if err == nil {
		t.Error("Expects Queue underflow error")
	}
}
