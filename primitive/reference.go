package primitive

// reference type is pointers in many programming languages
// pointers hold the memory address of a variable

type exampleStudentRegister struct {
	rollno      int64
	daysPresent int64
	daysAbsent  int64
}

var ptr = &exampleStudentRegister{}
