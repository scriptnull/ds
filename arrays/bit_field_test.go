package arrays

import "testing"

func TestBitField(t *testing.T) {
	var x BitField
	if x != BitField(0) {
		t.Error("Expected 0, but got", x)
	}

	x = x | 1
	if x != BitField(1) {
		t.Error("Expected 1, but got", x)
	}

	x = x & 1
	if x != BitField(1) {
		t.Error("Expected 1, but got", x)
	}
}
