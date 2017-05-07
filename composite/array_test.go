package composite

import (
	"testing"
)

func TestArrLength(t *testing.T) {
	if len(arr) != 5 {
		t.Error("Expected 5, got ", len(arr))
	}
}

func TestArrInitValues(t *testing.T) {
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			t.Error("Expected 0, got ", arr[i])
		}
	}
}

func TestColorsArray(t *testing.T) {
	if len(colors) != 5 {
		t.Error("Expected 5, got ", len(colors))
	}

	var expectedColors = [5]string{"green", "red", "blue", "cyan", "black"}

	for i := 0; i < len(colors); i++ {
		if colors[i] != expectedColors[i] {
			t.Errorf("Expected %s, got %s", expectedColors[i], colors[i])
		}
	}
}

func TestMatrixLength(t *testing.T) {
	if len(matrix) != 5 {
		t.Error("Expected 5, got ", len(matrix))
	}

	// matrix is fixed size (5) int8 array of fixed size (5) int8 arrays
}

func TestMatrixGetSet(t *testing.T) {
	var num int8 = 1
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			matrix[i][j] = num
			num++
		}
	}

	var expectedNum int8 = 1
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matrix[i][j] != expectedNum {
				t.Errorf("Expected %d, got %d", expectedNum, matrix[i][j])
			}
			expectedNum++
		}
	}
}

func TestJagged(t *testing.T) {
	if len(jagged) != 2 {
		t.Error("Expected 2, got ", len(jagged))
	}

	jagged[0] = []int8{21, 22, 12}
	jagged[1] = []int8{30, 1, 15, 22, 12}

	if len(jagged[0]) == len(jagged[1]) {
		t.Error("Expected variable length arrays to show example of jagged arrays")
	}
}
