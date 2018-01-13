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
		common.ColourWhite == pc.Colour && common.PieceKing == pc.Type {
		p.Board[common.Row8][common.ColumnD] = Piece{Colour: common.ColourWhite, Type: common.PieceRook}
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

	if isPieceCaptured {
		p.FiftyMoveClock = 0
	} else {
		p.FiftyMoveClock++
	}

	if colour == common.ColourBlack { // Increment clock if the movement is commited by black
		p.FullMoveClock++
	}

	return p
}
