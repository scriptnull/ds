package primitive

import (
	"reflect"
	"testing"
)

func TestPointerSize(t *testing.T) {
	valueSize := reflect.TypeOf(exampleStudentRegister{}).Size()
	ptrSize := reflect.TypeOf(ptr).Size()

	if valueSize != 24 {
		t.Error("Expected 24, got ", valueSize)
	}

	if ptrSize != 8 {
		t.Error("Expected 8, got ", ptrSize)
	}

	// Note that, pointer size is less than the actual datum size
	// So, pointer could be used for accessing the value efficiently
}
