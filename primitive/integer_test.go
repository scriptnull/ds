package primitive

import (
	"math"
	"reflect"
	"testing"
)

type testSizePair struct {
	variable interface{}
	// reflect.TypeOf().Size() returns size in bytes, so uint8 is enough to hold size values
	expectedSize uint8
}

func TestSignedIntSize(t *testing.T) {
	varIntSize := reflect.TypeOf(varInt).Size()

	if !((varIntSize == 4) || (varIntSize == 8)) {
		t.Error("Expected 4 or 8, got ", varIntSize)
	}

	var tests = []testSizePair{
		testSizePair{variable: varInt8, expectedSize: 1},
		testSizePair{variable: varInt16, expectedSize: 2},
		testSizePair{variable: varInt32, expectedSize: 4},
		testSizePair{variable: varInt64, expectedSize: 8},
	}

	for _, test := range tests {
		actualSize := uint8(reflect.TypeOf(test.variable).Size())
		if test.expectedSize != actualSize {
			t.Errorf("expected %d, got %d", test.expectedSize, actualSize)
		}
	}
}

func TestUnsignedIntSize(t *testing.T) {
	varUIntSize := reflect.TypeOf(varUInt).Size()

	if !((varUIntSize == 4) || (varUIntSize == 8)) {
		t.Error("Expected 4 or 8, got ", varUIntSize)
	}

	var tests = []testSizePair{
		testSizePair{variable: varUInt8, expectedSize: 1},
		testSizePair{variable: varUInt16, expectedSize: 2},
		testSizePair{variable: varUInt32, expectedSize: 4},
		testSizePair{variable: varUInt64, expectedSize: 8},
	}

	for _, test := range tests {
		actualSize := uint8(reflect.TypeOf(test.variable).Size())
		if test.expectedSize != actualSize {
			t.Errorf("expected %d, got %d", test.expectedSize, actualSize)
		}
	}
}

func TestSignedIntRange(t *testing.T) {
	maxInt8 := int8(math.Pow(2, 8)/2) - 1
	minInt8 := -1 * int8(math.Pow(2, 8)/2)

	if minInt8 != math.MinInt8 {
		t.Errorf("expected %d, got %d", math.MinInt8, minInt8)
	}

	if maxInt8 != math.MaxInt8 {
		t.Errorf("expected %d, got %d", math.MaxInt8, maxInt8)
	}

	// Min Max calculation follows the same pattern for other signed ints
}

func TestSignedInOverflow(t *testing.T) {
	// if tried to assign value, out of the accepted value, overflow happens
	varInt8 = int8(math.Pow(2, 8) / 2)

	if varInt8 != math.MinInt8 {
		t.Errorf("expected %d, got %d", math.MinInt8, varInt8)
	}

	// same pattern follows for other signed ints
}

func TestUnsignedIntRange(t *testing.T) {
	maxUInt8 := uint8(math.Pow(2, 8)) - 1

	// min value is always zero for all unsigned ints

	if maxUInt8 != math.MaxUint8 {
		t.Errorf("expected %d, got %d", math.MaxUint8, maxUInt8)
	}

	// Min Max calculation follows the same pattern for other unsigned ints
}

func TestUnsignedInOverflow(t *testing.T) {
	// if tried to assign value, out of the accepted value, overflow happens
	varUInt8 = uint8(math.Pow(2, 8))

	if varUInt8 != 0 {
		t.Errorf("expected %d, got %d", 0, varUInt8)
	}

	// same pattern follows for other unsigned ints
}
