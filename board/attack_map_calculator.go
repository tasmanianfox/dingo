package board

import "github.com/tasmanianfox/dingo/common"

// GetAttackMap returns a 2-dimensional array which represents a set of fields
// If specific fields is under attack of player c, the corresponding item equals "true"
func GetAttackMap(p Position, c int) [common.NumRows][common.NumColumns]bool {
	m := initAttackMap()
	for row, cells := range p.Board {
		for col, piece := range cells {
			if piece.Colour != c {
				continue
			}

			switch piece.Type {
			case common.PieceKing:
				m = mergeAttackMaps(m, getKingAttackMap(row, col))
			case common.PieceQueen:
				m = mergeAttackMaps(m, getQueenAttackMap(p.Board, row, col))
			case common.PieceRook:
				m = mergeAttackMaps(m, getRookAttackMap(p.Board, row, col))
			case common.PieceBishop:
				m = mergeAttackMaps(m, getBishopAttackMap(p.Board, row, col))
			case common.PieceKnight:
				m = mergeAttackMaps(m, getKnightAttackMap(row, col))
			case common.PiecePawn:
				m = mergeAttackMaps(m, getPawnAttackMap(row, col, p.ActiveColour))
			}
		}
	}

	return m
}

// IsKingChecked returns true if player's c king is checked at position p
func IsKingChecked(p Position, c int) bool {
	am := GetAttackMap(p, common.GetOpponent(c))
	r, c := p.LocateKing(p.ActiveColour)
	result := false
	if am[r][c] {
		result = true
	}

	return result
}

func getKingAttackMap(row int, col int) [common.NumRows][common.NumColumns]bool {
	m := initAttackMap()

	for sRow := row - 1; sRow <= row+1; sRow++ {
		if IsRowOutOfBoard(sRow) {
			continue
		}
		for sCol := col - 1; sCol <= col+1; sCol++ {
			if IsColumnOutOfBoard(sCol) || (sRow == row && sCol == col) {
				continue
			}

			m[sRow][sCol] = true
		}
	}

	return m
}

func getQueenAttackMap(b [common.NumRows][common.NumColumns]Piece, row int, col int) [common.NumRows][common.NumColumns]bool {
	return getVectorBasedAttackMap(
		[][2]int{{-1, 1}, {0, 1}, {1, 1}, {-1, 0}, {1, 0}, {-1, -1}, {0, -1}, {1, -1}},
		b,
		row,
		col,
	)
}

func getRookAttackMap(b [common.NumRows][common.NumColumns]Piece, row int, col int) [common.NumRows][common.NumColumns]bool {
	return getVectorBasedAttackMap(
		[][2]int{{0, 1}, {-1, 0}, {1, 0}, {0, -1}},
		b,
		row,
		col,
	)
}

func getBishopAttackMap(b [common.NumRows][common.NumColumns]Piece, row int, col int) [common.NumRows][common.NumColumns]bool {
	return getVectorBasedAttackMap(
		[][2]int{{-1, 1}, {1, 1}, {-1, -1}, {1, -1}},
		b,
		row,
		col,
	)
}

func getKnightAttackMap(row int, col int) [common.NumRows][common.NumColumns]bool {
	m := initAttackMap()
	os := [][2]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}} // offsets
	for _, o := range os {
		testRow := row + o[0]
		testCol := col + o[1]
		if IsColumnOutOfBoard(testRow) || IsRowOutOfBoard(testCol) {
			continue
		}

		m[testRow][testCol] = true
	}

	return m
}

func getPawnAttackMap(row int, col int, colour int) [common.NumRows][common.NumColumns]bool {
	m := initAttackMap()
	testRow := 1
	if common.ColourBlack == colour {
		testRow = -1
	}
	testRow = row + testRow
	if !IsRowOutOfBoard(testRow) {
		for _, colOffset := range [2]int{-1, 1} {
			testCol := col + colOffset
			if IsColumnOutOfBoard(testCol) {
				continue
			}
			m[testCol][testRow] = true
		}
	}

	return m
}

// getVectorBasedAttackMap calculates attack map for pieces with linear attacks (queen, rook, bishop)
func getVectorBasedAttackMap(vs [][2]int, b [common.NumRows][common.NumColumns]Piece, row int, col int) [common.NumRows][common.NumColumns]bool {
	m := [common.NumRows][common.NumColumns]bool{}

	for _, vector := range vs {
		rowVector := vector[0]
		colVector := vector[1]

		for lineIndex := 1; ; lineIndex++ {
			testRow := lineIndex*rowVector + row
			testColumn := lineIndex*colVector + col
			if IsColumnOutOfBoard(testColumn) || IsRowOutOfBoard(testRow) {
				break
			}

			m[testRow][testColumn] = true

			if b[testRow][testColumn].Type != common.PieceEmpty {
				break
			}
		}
	}

	return m
}

func initAttackMap() [common.NumRows][common.NumColumns]bool {
	m := [common.NumRows][common.NumColumns]bool{}
	for row := 0; row < common.NumRows; row++ {
		for col := 0; row < common.NumColumns; col++ {
			m[row][col] = false
		}
	}

	return m
}

func mergeAttackMaps(a [common.NumRows][common.NumColumns]bool, b [common.NumRows][common.NumColumns]bool) [common.NumRows][common.NumColumns]bool {
	c := [common.NumRows][common.NumColumns]bool{}
	for row := 0; row < common.NumRows; row++ {
		for col := 0; col < common.NumColumns; col++ {
			c[row][col] = a[col][row] || b[col][row]
		}
	}

	return c
}
