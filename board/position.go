package board

import "github.com/tasmanianfox/dingo/common"

type Position struct {
	// The first dimension is a row, the second is a column
	Board                  [8][8]Piece
	ActiveColour           int
	WhiteKingsideCastling  bool
	WhiteQueensideCastling bool
	BlackKingsideCastling  bool
	BlackQueensideCastling bool
	// -1 if no target. Points to pawn (Row 4 or 5)
	EnPassantTargetRow int
	// -1 if no target
	EnPassantTargetColumn int
	FiftyMoveClock        int
	FullMoveClock         int
}

// IsEmptyCell returns true if the cell at column c and row r is empty
func (p Position) IsEmptyCell(c int, r int) bool {
	return common.ColourEmpty == p.Board[r][c].Colour && common.PieceEmpty == p.Board[r][c].Type
}

// GetPieceAt returns a piece that is located at column c and row r
func (p Position) GetPieceAt(c int, r int) Piece {
	return p.Board[r][c]
}

// GetEmptyCell returns an "empty" piece which indicates that the cell is empty
func GetEmptyCell() Piece {
	return Piece{Colour: common.ColourEmpty, Type: common.PieceEmpty}
}
