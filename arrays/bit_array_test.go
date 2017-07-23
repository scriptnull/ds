package arrays

import "testing"

func TestNewBitArray(t *testing.T) {
	cases := []struct {
		input    int
		expected int
	}{
		{1, 1}, {10, 2}, {32, 4},
	}

	checkSize := func(size, expected int) {
		a := NewBitArray(size)
		bArray := []int8(a)
		if len(bArray) != expected {
			t.Errorf("Expected length to be %d, but got %d", expected, len(bArray))
		}
	}

	for _, testCase := range cases {
		checkSize(testCase.input, testCase.expected)
	}
}
