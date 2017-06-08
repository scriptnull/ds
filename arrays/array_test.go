package arrays

import (
	"strconv"
	"testing"
)

func TestSingleDimensionalArray(t *testing.T) {
	if arr[0] != 1 {
		t.Error("Expected 1, got ", arr[0])
	}

	for index, expected := 0, 1; index < len(arr); index, expected = index+1, expected+1 {
		if arr[index] != expected {
			t.Errorf("Expected %d, got %d", expected, arr[index])
		}
	}
}

func TestMultiDimensionalArray(t *testing.T) {
	for rowIndex, row := range multi {
		for colIndex := 0; colIndex < len(row); colIndex++ {
			expected, err := strconv.Atoi(strconv.Itoa(rowIndex) + strconv.Itoa(colIndex))
			if err != nil {
				t.Error("Expected err to be not null, but got ", err)
			}
			if multi[rowIndex][colIndex] != expected {
				t.Errorf("Expected %d, got %d", expected, multi[rowIndex][colIndex])
			}
		}
	}
}
