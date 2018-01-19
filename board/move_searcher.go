package board

import (
	"github.com/tasmanianfox/dingo/common"
)

type castlingData struct {
	Colour        int
	KingsideFlag  bool
	QueensideFlag bool
	Row           int
}

type pawnData struct {
	Colour    int
	FistRow   int
	Increment int
}

// FindAllAvailableMoves finds all available moves for active player
func FindAllAvailableMoves(p Position) []Move {
	ms := []Move{}
	for row, cells := range p.Board {
		for col, piece := range cells {
			if piece.Colour != p.ActiveColour {
				continue
			}

			switch piece.Type {
			case common.PieceKing:
				ms = append(ms, findKingMoves(p, row, col)...)
			case common.PieceQueen:
				ms = append(ms, findUsualPieceMoves(p, row, col)...)
			case common.PieceRook:
				ms = append(ms, findUsualPieceMoves(p, row, col)...)
			case common.PieceBishop:
				ms = append(ms, findUsualPieceMoves(p, row, col)...)
			case common.PieceKnight:
				ms = append(ms, findUsualPieceMoves(p, row, col)...)
			case common.PiecePawn:
				ms = append(ms, findPawnMoves(p, row, col)...)
			}
		}
	}

	return ms
}

func findKingMoves(p Position, row int, col int) []Move {
	ms := []Move{}
	oam := GetAttackMap(p, common.GetOpponent(p.ActiveColour))

	for sRow := row - 1; sRow <= row+1; sRow++ {
		if IsRowOutOfBoard(sRow) {
			continue
		}
		for sCol := col + 1; sCol >= col-1; sCol-- {
			if IsColumnOutOfBoard(sCol) || (sRow == row && sCol == col) || true == oam[sRow][sCol] || p.ActiveColour == p.Board[sRow][sCol].Colour {
				continue
			}

			ms = appendMoveIfValid(p, ms, Move{SourceRow: row, SourceColumn: col, DestRow: sRow, DestColumn: sCol})
		}
	}

	// Castlings
	cd := [2]castlingData{
		castlingData{Colour: common.ColourWhite, KingsideFlag: p.WhiteKingsideCastling, QueensideFlag: p.WhiteQueensideCastling, Row: common.Row1},
		castlingData{Colour: common.ColourBlack, KingsideFlag: p.BlackKingsideCastling, QueensideFlag: p.BlackQueensideCastling, Row: common.Row8},
	}

	for _, castlingDatum := range cd {
		if castlingDatum.Colour == p.ActiveColour && !oam[castlingDatum.Row][common.ColumnE] {
			if castlingDatum.KingsideFlag && !oam[castlingDatum.Row][common.ColumnF] &&
				p.IsEmptyCell(common.ColumnF, castlingDatum.Row) && p.IsEmptyCell(common.ColumnG, castlingDatum.Row) {
				ms = appendMoveIfValid(p, ms, Move{SourceColumn: common.ColumnE, SourceRow: castlingDatum.Row, DestRow: castlingDatum.Row, DestColumn: common.ColumnG})
			}
			if castlingDatum.QueensideFlag && !oam[castlingDatum.Row][common.ColumnD] && !oam[castlingDatum.Row][common.ColumnC] &&
				p.IsEmptyCell(common.ColumnB, castlingDatum.Row) && p.IsEmptyCell(common.ColumnC, castlingDatum.Row) &&
				p.IsEmptyCell(common.ColumnD, castlingDatum.Row) {
				ms = appendMoveIfValid(p, ms, Move{SourceColumn: common.ColumnE, SourceRow: castlingDatum.Row, DestRow: castlingDatum.Row, DestColumn: common.ColumnC})
			}
		}
	}

	return ms
}

func findUsualPieceMoves(p Position, row int, col int) []Move {
	ms := []Move{}
	am := [common.NumRows][common.NumColumns]bool{}
	switch p.Board[row][col].Type {
	case common.PieceKnight:
		am = getKnightAttackMap(row, col)
	case common.PieceBishop:
		am = getBishopAttackMap(p.Board, row, col)
	case common.PieceRook:
		am = getRookAttackMap(p.Board, row, col)
	case common.PieceQueen:
		am = getQueenAttackMap(p.Board, row, col)
	}

	for testRow, attackMapColumns := range am {
		for testCol, attackMapColumn := range attackMapColumns {
			if false == attackMapColumn || p.ActiveColour == p.Board[testRow][testCol].Colour {
				continue
			}
			ms = appendMoveIfValid(p, ms, Move{SourceRow: row, SourceColumn: col, DestRow: testRow, DestColumn: testCol})
		}
	}

	return ms
}

func findPawnMoves(p Position, row int, col int) []Move {
	ms := []Move{}

	// Moves
	pda := []pawnData{
		pawnData{Colour: common.ColourWhite, FistRow: common.Row2, Increment: 1},
		pawnData{Colour: common.ColourBlack, FistRow: common.Row7, Increment: -1},
	}

	for _, pd := range pda {
		if pd.Colour != p.ActiveColour {
			continue
		}
		destRow := row + pd.Increment
		if p.IsEmptyCell(col, destRow) {
			ms = appendMoveIfValid(p, ms, Move{SourceRow: row, SourceColumn: col, DestRow: destRow, DestColumn: col})
			if row == pd.FistRow {
				destRow := destRow + pd.Increment
				if p.IsEmptyCell(col, destRow) {
					ms = appendMoveIfValid(p, ms, Move{SourceRow: row, SourceColumn: col, DestRow: destRow, DestColumn: col})
				}
			}
		}
	}

	// Attacks
	am := getPawnAttackMap(row, col, p.ActiveColour)
	for testRow, cells := range am {
		for testCol, isUnderAttack := range cells {
			if !isUnderAttack {
				continue
			}

			if (p.Board[testRow][testCol].Colour == common.GetOpponent(p.ActiveColour)) ||
				(p.EnPassantTargetColumn == testCol && p.EnPassantTargetRow == row) {
				ms = appendMoveIfValid(p, ms, Move{SourceRow: row, SourceColumn: col, DestRow: testRow, DestColumn: testCol})
			}
		}
	}

	return ms
}

// Validates the new position. If position is valid appends the move to array
func appendMoveIfValid(p Position, ms []Move, m Move) []Move {
	p2 := CommitMove(p, m)
	if !p2.IsKingChecked(p.ActiveColour) {
		ms = append(ms, m)
	}

	return ms
}
