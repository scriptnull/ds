package primitive

import "testing"

func TestBoolean(t *testing.T) {
	if truthy != true {
		t.Error("Expected true, got ", truthy)
	}

	if falsy != false {
		t.Error("Expected false, got ", falsy)
	}

}
