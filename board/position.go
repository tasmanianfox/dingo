package board

type Position struct {
	// The first dimension is a row, the second is a column
	Board                  [8][8]Piece
	ActiveColour           int
	WhiteKingsideCastling  bool
	WhiteQueensideCastling bool
	BlackKingsideCastling  bool
	BlackQueensideCastling bool
	// -1 if no target
	EnPassantTargetRow int
	// -1 if no target
	EnPassantTargetColumn int
	FiftyMoveClock        int
	FullMoveClock         int
}
