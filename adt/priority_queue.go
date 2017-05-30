package adt

// PriorityQueue Specification
type PriorityQueue interface {
	AddElement(int) error
	// Remove element with highest or lowest priority from queue
	PullElement() (int, error)
	// Read the value of the highest or lowest priority in queue
	PeekElement() (int, error)
}
