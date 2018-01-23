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
	resultRow, resultCol := -1, -1
	for row, cells := range p.Board {
		for column, piece := range cells {
			if piece.Colour == c && piece.Type == common.PieceKing {
				resultRow = row
				resultCol = column
				break
			}
		}
	}

	if -1 == resultRow || -1 == resultCol {
		panic("Could not locate the king of player " + strconv.Itoa(c))
	}

	return resultRow, resultCol
}

// IsKingChecked returns true if player's c king is checked at position p
func (p Position) IsKingChecked(c int) bool {
	am := GetAttackMap(p, common.GetOpponent(c))
	row, col := p.LocateKing(c)
	result := false
	if am[row][col] {
		result = true
	}

	return result
}

// IsKingChecked returns true if player's c king is checkmnated at position p
func (p Position) IsKingCheckmated(c int) bool {
	p2 := p
	p2.ActiveColour = c
	result := false
	cf := p.IsKingChecked(c)
	if cf {
		ms := FindAllAvailableMoves(p2)
		if 0 == len(ms) {
			result = true
		}
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
