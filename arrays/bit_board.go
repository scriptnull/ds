package arrays

// BitBoard is a data structure commonly used in computer systems that play board games.
// A bitboard, often used for boardgames such as chess, checkers, othello and word games, is a specialization of the bit array data structure, where each bit represents a game position or state, designed for optimization of speed and/or memory or disk use in mass calculations.
type BitBoard interface {
	// AND all the bits from right to left in a row
	AndAllRightToLeft(rowIndex int)
	// AND all the bits from top to bottom in a column
	AndAllTopToBottom(columnIndex int)
	// OR all the bits from right to left in a row
	OrAllRightToLeft(rowIndex int)
	// OR all the bits from top to bottom in a column
	OrAllTopToBottom(columnIndex int)
	// Rotate the Board by 45 degrees
	RotateBy45()
	// Rotate the Board by 90 degrees
	RotateBy90()
}
