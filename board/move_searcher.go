package board

import (
	"github.com/tasmanianfox/dingo/common"
)

// FindAllAvailableMoves finds all available moves for active player
func FindAllAvailableMoves(p Position) []Move {
	ms := []Move{}
	for row, cells := range p.Board {
		for col, piece := range cells {
			if piece.Colour != p.ActiveColour {
				continue
			}

			ms2 := []Move{}
			switch piece.Type {
			case common.PieceKing:
				ms2 = append(ms, findKingMoves(p, row, col)...)
			case common.PieceQueen:
				ms2 = append(ms, findUsualPieceMoves(p, row, col)...)
			case common.PieceRook:
				ms2 = append(ms, findUsualPieceMoves(p, row, col)...)
			case common.PieceBishop:
				ms2 = append(ms, findUsualPieceMoves(p, row, col)...)
			case common.PieceKnight:
				ms2 = append(ms, findUsualPieceMoves(p, row, col)...)
			case common.PiecePawn:
				ms2 = append(ms, findPieceMoves(p, row, col)...)
			}

			// Filter out those moves that lead to own king being checked
			for _, m := range ms2 {
				p2 := CommitMove(p, m)
				if !p2.IsKingChecked(p.ActiveColour) {
					ms = append(ms, m)
				}
			}
		}
	}

	return ms
}

func findKingMoves(p Position, row int, col int) []Move {
	ms := []Move{}
	oc := common.GetOpponent(p.ActiveColour)
	oam := GetAttackMap(p, oc)

	for sRow := row - 1; sRow <= row+1; sRow++ {
		if IsRowOutOfBoard(sRow) {
			continue
		}
		for sCol := col - 1; sCol <= col+1; sCol++ {
			if IsColumnOutOfBoard(sCol) || (sRow == row && sCol == col) || true == oam[sRow][sCol] || p.ActiveColour == p.Board[sRow][sCol].Colour {
				continue
			}

			ms = append(ms, Move{SourceRow: row, SourceColumn: col, DestRow: sRow, DestColumn: sCol})
		}
	}

	// Castlings
	if common.ColourWhite == p.ActiveColour && !oam[common.Row1][common.ColumnE] {
		if p.WhiteKingsideCastling && !oam[common.Row1][common.ColumnF] {
			ms = append(ms, Move{DestRow: common.Row1, DestColumn: common.ColumnG})
		}
		if p.WhiteQueensideCastling && !oam[common.Row1][common.ColumnD] && !oam[common.Row1][common.ColumnC] {
			ms = append(ms, Move{DestRow: common.Row1, DestColumn: common.ColumnC})
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

	for testRow, cells := range am {
		for testCol := range cells {
			if p.ActiveColour == p.Board[testRow][testRow].Colour {
				continue
			}
			ms = append(ms, Move{SourceRow: row, SourceColumn: col, DestRow: testRow, DestColumn: testCol})
		}
	}

	return ms
}

func findPieceMoves(p Position, row int, col int) []Move {
	ms := []Move{}

	am := getPawnAttackMap(row, col, p.ActiveColour)
	for testRow, cells := range am {
		for testCol, isUnderAttack := range cells {
			if !isUnderAttack {
				continue
			}

			if (p.Board[testRow][testCol].Colour == common.GetOpponent(p.ActiveColour)) ||
				(p.EnPassantTargetColumn == testCol && p.EnPassantTargetRow == row) {
				ms = append(ms, Move{SourceRow: row, SourceColumn: col, DestRow: testRow, DestColumn: testCol})
			}
		}
	}

	return ms
}
