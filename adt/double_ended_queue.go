package adt

// DoubleEndedQueue is a specification of compulsory functions for Double ended queue implementation
type DoubleEndedQueue interface {
	Push(int) error
	Pop() (int, error)
	UnShift(int) error
	Shift() (int, error)
}
