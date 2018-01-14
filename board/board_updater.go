package board

import (
	"strconv"

	"github.com/tasmanianfox/dingo/common"
)

// CommitMovement Performs a movement m using position p
func CommitMovement(p Position, m Movement) Position {
	pc := p.Board[m.SourceRow][m.SourceColumn]
	if pc.Colour == common.ColourEmpty || pc.Type == common.PieceEmpty {
		panic("Cannot move from specified cell. The cell is empty: type " + strconv.Itoa(pc.Type) + " / colour " +
			strconv.Itoa(pc.Colour) + ", position: " + PositionToFen(p) + ", movement " + strconv.Itoa(m.SourceColumn) + strconv.Itoa(m.SourceRow) + "-" +
			strconv.Itoa(m.DestColumn) + strconv.Itoa(m.DestRow))
	}
	isPieceCaptured := !p.IsEmptyCell(m.DestColumn, m.DestRow)
	colour := p.GetPieceAt(m.SourceColumn, m.SourceRow).Colour

	p.Board[m.SourceRow][m.SourceColumn] = GetEmptyCell()
	p.Board[m.DestRow][m.DestColumn] = pc

	if isPieceCaptured || common.PiecePawn == pc.Type {
		p.FiftyMoveClock = 0
	} else {
		p.FiftyMoveClock++
	}

	if colour == common.ColourBlack { // Increment clock if the movement is commited by black
		p.FullMoveClock++
	}

	// Castlings
	// White kingside
	if common.Row1 == m.SourceRow && common.Row1 == m.DestRow &&
		common.ColumnE == m.SourceColumn && common.ColumnG == m.DestColumn &&
		common.ColourWhite == pc.Colour && common.PieceKing == pc.Type {
		p.Board[common.Row1][common.ColumnF] = Piece{Colour: common.ColourWhite, Type: common.PieceRook}
		p.Board[common.Row1][common.ColumnH] = GetEmptyCell()
	}
	// White queenside
	if common.Row1 == m.SourceRow && common.Row1 == m.DestRow &&
		common.ColumnE == m.SourceColumn && common.ColumnC == m.DestColumn &&
		common.ColourWhite == pc.Colour && common.PieceKing == pc.Type {
		p.Board[common.Row1][common.ColumnD] = Piece{Colour: common.ColourWhite, Type: common.PieceRook}
		p.Board[common.Row1][common.ColumnA] = GetEmptyCell()
	}
	// Black kingside
	if common.Row8 == m.SourceRow && common.Row8 == m.DestRow &&
		common.ColumnE == m.SourceColumn && common.ColumnG == m.DestColumn &&
		common.ColourBlack == pc.Colour && common.PieceKing == pc.Type {
		p.Board[common.Row8][common.ColumnF] = Piece{Colour: common.ColourBlack, Type: common.PieceRook}
		p.Board[common.Row8][common.ColumnH] = GetEmptyCell()
	}
	// Black queenside
	if common.Row8 == m.SourceRow && common.Row8 == m.DestRow &&
		common.ColumnE == m.SourceColumn && common.ColumnC == m.DestColumn &&
		common.ColourBlack == pc.Colour && common.PieceKing == pc.Type {
		p.Board[common.Row8][common.ColumnD] = Piece{Colour: common.ColourBlack, Type: common.PieceRook}
		p.Board[common.Row8][common.ColumnA] = GetEmptyCell()
	}

	// En passant
	if p.EnPassantTargetColumn == m.DestColumn && common.ColourWhite == pc.Colour && p.EnPassantTargetRow == m.DestRow-1 {
		p.Board[m.DestRow-1][m.DestColumn] = GetEmptyCell()
		isPieceCaptured = true
	}
	if p.EnPassantTargetColumn == m.DestColumn && common.ColourBlack == pc.Colour && p.EnPassantTargetRow == m.DestRow+1 {
		p.Board[m.DestRow+1][m.DestColumn] = GetEmptyCell()
		isPieceCaptured = true
	}

	// Promotion
	if common.Row8 == m.DestRow && common.PiecePawn == pc.Type && common.ColourWhite == pc.Colour {
		p.Board[m.DestRow][m.DestColumn].Type = m.CastTo
	}
	if common.Row1 == m.DestRow && common.PiecePawn == pc.Type && common.ColourBlack == pc.Colour {
		p.Board[m.DestRow][m.DestColumn].Type = m.CastTo
	}

	// Disable castling
	if common.Row1 == m.DestRow {
		if true == p.WhiteKingsideCastling && (common.ColumnH == m.SourceColumn || common.ColumnH == m.DestColumn || common.ColumnE == m.SourceColumn) {
			p.WhiteKingsideCastling = false
		}
		if true == p.WhiteQueensideCastling && (common.ColumnA == m.SourceColumn || common.ColumnA == m.DestColumn || common.ColumnE == m.SourceColumn) {
			p.WhiteQueensideCastling = false
		}
	}
	if common.Row8 == m.DestRow {
		if true == p.BlackKingsideCastling && (common.ColumnH == m.SourceColumn || common.ColumnH == m.DestColumn || common.ColumnE == m.SourceColumn) {
			p.BlackKingsideCastling = false
		}
		if true == p.BlackQueensideCastling && (common.ColumnA == m.SourceColumn || common.ColumnA == m.DestColumn || common.ColumnE == m.SourceColumn) {
			p.BlackQueensideCastling = false
		}
	}

	// Assign the en passant target
	if (common.ColourWhite == colour && common.Row2 == m.SourceRow && common.Row4 == m.DestRow) ||
		common.ColourBlack == colour && common.Row7 == m.SourceRow && common.Row5 == m.DestRow {
		p.EnPassantTargetColumn = m.DestColumn
		p.EnPassantTargetRow = m.DestRow
	} else {
		p.EnPassantTargetColumn = -1
		p.EnPassantTargetRow = -1
	}

	// Change active player
	if colour == common.ColourWhite {
		p.ActiveColour = common.ColourBlack
	} else {
		p.ActiveColour = common.ColourWhite
	}

	return p
}
