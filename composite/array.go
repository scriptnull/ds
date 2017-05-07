package composite

// Basic array - fixed size of 5. Index starts from 0 and ends at 4
var arr [5]int8

// Initialized array
var colors = [5]string{"green", "red", "blue", "cyan", "black"}

// MultiDimensional arrays
var matrix [5][5]int8

// Initialized MultiDimensional array
// Imagining, two row of seat in a uber ride can be expressed as a multi dimensional array
var uberPool = [2][2]bool{
	[2]bool{false, true},
	[2]bool{true, true},
}

// Jagged arrays - arrays having rows of varying length
// Submitting notebooks, assignment in a class is actually a jagged array
// Each notebook that are submitted are arranged sequentially on its own collection
// Each notebook category, does not need to have the same number of notes ( may be some student forget to bring his homework note)
var jagged [2][]int8
