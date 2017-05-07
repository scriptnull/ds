package primitive

// ByteSize is an enum type of general memory representations
type ByteSize float64

const (
	_ = iota
	// KB is kilobytes of memory
	KB ByteSize = 1 << (10 * iota)
	// MB is MegaBytes of memory
	MB
)

// Color is an emum type of colors
// It represents
type Color uint8

const (
	// C1 is first 8 bit color
	C1 Color = iota
	// C2 is second 8 bit color
	C2
)
