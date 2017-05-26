package adt

import "errors"

// ArrStack is sized array implementation
type ArrStack struct {
	arr [5]int
	top int
}

// Push an element to the stack
func (stack *ArrStack) Push(element int) error {
	if stack.top == len(stack.arr) {
		return errors.New("Stack overflow")
	}

	stack.arr[stack.top] = element
	stack.top++

	return nil
}

// Pop and element from the stack
func (stack *ArrStack) Pop() (int, error) {
	if stack.top == 0 {
		return 0, errors.New("Stack underflow")
	}

	topElement := stack.arr[stack.top-1]
	stack.top--

	return topElement, nil
}

// Peek at the top element of the stack
func (stack *ArrStack) Peek() int {
	return stack.arr[stack.top]
}

// DynStack is dynamic array implementation
type DynStack struct {
	arr []int
	top int
}

// Push an element to the stack
func (stack *DynStack) Push(element int) error {
	stack.arr = append(stack.arr, element)
	stack.top++

	return nil
}

// Pop and element from the stack
func (stack *DynStack) Pop() (int, error) {
	if stack.top == 0 {
		return 0, errors.New("Stack underflow")
	}

	topElement := stack.arr[stack.top-1]
	stack.top--

	return topElement, nil
}

// Peek at the top element of the stack
func (stack *DynStack) Peek() int {
	if len(stack.arr) == 0 {
		return 0
	}
	return stack.arr[stack.top]
}

// LinkedStack is a Linked List implementation of stack
type LinkedStack struct {
	top  *stackItem
	size int
}
type stackItem struct {
	value int
	next  *stackItem
}

// Push an element to the stack
func (stack *LinkedStack) Push(element int) error {
	newItem := &stackItem{
		value: element,
		next:  stack.top,
	}

	stack.top = newItem
	stack.size++

	return nil
}

// Pop and element from the stack
func (stack *LinkedStack) Pop() (int, error) {
	topItem := stack.top
	if topItem == nil {
		return 0, errors.New("Stack underflow")
	}

	stack.top = stack.top.next

	return topItem.value, nil
}

// Peek at the top element of the stack
func (stack *LinkedStack) Peek() int {
	if stack.top == nil {
		return 0
	}
	return stack.top.value
}
