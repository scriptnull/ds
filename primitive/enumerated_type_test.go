package primitive

import (
	"testing"
)

func TestEnumValues(t *testing.T) {
	// enums can be compared with its actual value
	if C1 != 0 {
		t.Error("Expected 0, got ", C1)
	}

	// enums cannot be compared with other type of enums
	// Example: C1 == KB will result in compile time error
	// invalid operation: C1 == KB (mismatched types Color and ByteSize)
}
