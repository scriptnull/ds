package composite

// record is collection of fields possibly of different datatype
// records are implemented as structs in golang

// Employee is a sample record
type Employee struct {
	ID        uint8
	FirstName string
	LastName  string
}
