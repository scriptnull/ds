package primitive

import (
	"reflect"
	"testing"
)

func TestFloatSize(t *testing.T) {
	actualSize := reflect.TypeOf(varFloat32).Size()

	if actualSize != 4 {
		t.Error("Expected 4, got ", actualSize)
	}
}
