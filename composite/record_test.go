package composite

import "testing"

func TestAssignment(t *testing.T) {
	var emp1 = &Employee{
		ID:        1,
		FirstName: "Jake",
		LastName:  "Long",
	}

	if emp1 == nil {
		t.Error("Expected non empty value")
	}
}

func TestAccess(t *testing.T) {
	var emp1 = &Employee{
		ID:        1,
		FirstName: "Jake",
		LastName:  "Long",
	}

	if emp1.ID != 1 {
		t.Error("Expected 1, got ", emp1.ID)
	}
	if emp1.FirstName != "Jake" {
		t.Error("Expected Jake, got ", emp1.FirstName)
	}
	if emp1.LastName != "Long" {
		t.Error("Expected Long, got ", emp1.LastName)
	}
}
