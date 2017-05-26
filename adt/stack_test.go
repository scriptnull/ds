package adt

import "testing"

func TestSizedArrayImplementation(t *testing.T) {
	var stack = ArrStack{}

	_, err := stack.Pop()
	if err == nil {
		t.Error("Expected Stack underflow error")
	}

	ele := stack.Peek()
	if ele != 0 {
		t.Error("Expected 0, got ", ele)
	}

	err = stack.Push(1)
	if err != nil {
		t.Error("Expected to push element onto stack, but got ", err)
	}
	_ = stack.Push(2)
	_ = stack.Push(3)
	_ = stack.Push(4)
	_ = stack.Push(5)
	err = stack.Push(6)
	if err == nil {
		t.Error("Expected Stack overflow error")
	}
}

func TestDynamicArrayImplementation(t *testing.T) {
	var stack = DynStack{}

	_, err := stack.Pop()
	if err == nil {
		t.Error("Expected Stack underflow error")
	}

	ele := stack.Peek()
	if ele != 0 {
		t.Error("Expected 0, got ", ele)
	}

	err = stack.Push(1)
	if err != nil {
		t.Error("Expected to push element onto stack, but got ", err)
	}
	_ = stack.Push(2)
	_ = stack.Push(3)
	_ = stack.Push(4)
	_ = stack.Push(5)
	err = stack.Push(6)
	if err != nil {
		t.Error("Expected Stack overflow error")
	}
}

func TestLinkedListImplementation(t *testing.T) {
	var stack = LinkedStack{}

	_, err := stack.Pop()
	if err == nil {
		t.Error("Expected Stack underflow error")
	}

	ele := stack.Peek()
	if ele != 0 {
		t.Error("Expected 0, got ", ele)
	}

	err = stack.Push(1)
	if err != nil {
		t.Error("Expected to push element onto stack, but got ", err)
	}
	_ = stack.Push(2)
	_ = stack.Push(3)
	_ = stack.Push(4)
	_ = stack.Push(5)
	err = stack.Push(6)
	if err != nil {
		t.Error("Expected Stack overflow error")
	}
}
