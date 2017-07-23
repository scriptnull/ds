package arrays

// BitArray is used to implement Set ADT
// Also known as Bit map, Bit set, Bit vector
// A typical bit array stores kw bits, where w is the number of bits in the unit of storage, such as a byte or word, and k is some nonnegative integer.
// If w does not divide the number of bits to be stored, some space is wasted due to internal fragmentation.
type BitArray []int8

// NewBitArray initializes a new bitarray with k*w bit array store
// w is number if bits in unit storage, such as byte or word => in this case its int8 i.e a byte
// k is the size argument
func NewBitArray(size int) BitArray {
	numOfByteAllocs := size / 8
	if size%8 != 0 {
		return BitArray(make([]int8, numOfByteAllocs+1, numOfByteAllocs+1))
	}
	return BitArray(make([]int8, numOfByteAllocs, numOfByteAllocs))
}

// Practical application include having a BitArray of size 31
// to represent the attendance of student, with 1 or 0 as present and absent to class

// Checking i'th is set is one common operation, that has to be supported
