package arrays

import "testing"

func TestArray(t *testing.T) {
	if arr[0] != 1 {
		t.Error("Expected 1, got ", arr[0])
	}

	for index, expected := 0, 1; index < len(arr); index, expected = index+1, expected+1 {
		if arr[index] != expected {
			t.Errorf("Expected %d, got %d", expected, arr[index])
		}
	}
}
