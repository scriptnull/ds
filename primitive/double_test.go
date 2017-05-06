package primitive

import (
	"reflect"
	"testing"
)

func TestDoubleSize(t *testing.T) {
	actualSize := reflect.TypeOf(varFloat64).Size()

	if actualSize != 8 {
		t.Error("Expected 4, got ", actualSize)
	}
}
