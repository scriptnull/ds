package adt

// DoubleEndedQueue is a specification of compulsory functions for Double ended queue implementation
type DoubleEndedQueue interface {
	Push(int)
	Pop() int
	UnShift(int)
	Shift() int
}
