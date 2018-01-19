package board

import (
	"strconv"

	"github.com/tasmanianfox/dingo/common"
)

// Position describes current position (piece alignment, availability of en passant and castlings, current move number, etc)
type Position struct {
	// The first dimension is a row, the second is a column
	Board                  [common.NumRows][common.NumColumns]Piece
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
	return common.ColourEmpty == p.Board[r][c].Colour
}

// GetPieceAt returns a piece that is located at column c and row r
func (p Position) GetPieceAt(c int, r int) Piece {
	return p.Board[r][c]
}

// LocateKing returns row and column of player's c king
func (p Position) LocateKing(c int) (int, int) {
	r, c := -1, -1
	for row, cells := range p.Board {
		for column, piece := range cells {
			if piece.Colour == c && piece.Type == common.PieceKing {
				r = row
				c = column
				break
			}
		}
	}

	if -1 == r || -1 == c {
		panic("Could not locate the king of player " + strconv.Itoa(c))
	}

	return r, c
}

// IsKingChecked returns true if player's c king is checked at position p
func (p Position) IsKingChecked(c int) bool {
	am := GetAttackMap(p, common.GetOpponent(c))
	r, c := p.LocateKing(p.ActiveColour)
	result := false
	if am[r][c] {
		result = true
	}

	return result
}

// GetEmptyCell returns an "empty" piece which indicates that the cell is empty
func GetEmptyCell() Piece {
	return Piece{Colour: common.ColourEmpty, Type: common.PieceEmpty}
}

// IsRowOutOfBoard returns true if the specified row index is out of board dimension
func IsRowOutOfBoard(r int) bool {
	return r < 0 || r >= common.NumRows
}

// IsColumnOutOfBoard returns true if the specified column index is out of board dimension
func IsColumnOutOfBoard(c int) bool {
	return c < 0 || c >= common.NumColumns
}
