package board

import (
	"strconv"
	"strings"

	"github.com/tasmanianfox/dingo/common"
)

// FenToPosition converts FEN representation of the position into board.Position object
func FenToPosition(fen string) Position {
	var data = strings.Split(fen, " ")
	var pos Position

	// Set pieces
	var rows = strings.Split(data[0], "/")
	for invRowIndex, row := range rows {
		var rowIndex = common.NumRows - invRowIndex - 1
		var colIndex = 0
		for colIndex < len(row) {
			var cell = row[colIndex : colIndex+1]
			var pc Piece
			// Cell is numeric: create empty cells
			if numEmptyCells, err := strconv.Atoi(cell); err == nil {
				pc = Piece{Colour: common.ColourEmpty, Type: common.PieceEmpty}
				for i := 0; i < numEmptyCells; i++ {
					pos.Board[rowIndex][colIndex] = pc
					colIndex++
				}
			} else {
				if strings.ToUpper(cell) == cell {
					pc.Colour = common.ColourWhite
				} else {
					pc.Colour = common.ColourBlack
				}
				pc.Type = charToPiece(strings.ToLower(cell))
				pos.Board[rowIndex][colIndex] = pc
				colIndex++
			}
		}
	}

	// Set active colour
	if "w" == data[1] {
		pos.ActiveColour = common.ColourWhite
	} else {
		pos.ActiveColour = common.ColourBlack
	}

	// Castling
	if strings.Index(data[2], "K") > -1 {
		pos.WhiteKingsideCastling = true
	} else {
		pos.WhiteKingsideCastling = false
	}
	if strings.Index(data[2], "Q") > -1 {
		pos.WhiteQueensideCastling = true
	} else {
		pos.WhiteQueensideCastling = false
	}
	if strings.Index(data[2], "k") > -1 {
		pos.BlackKingsideCastling = true
	} else {
		pos.BlackKingsideCastling = false
	}
	if strings.Index(data[2], "q") > -1 {
		pos.BlackQueensideCastling = true
	} else {
		pos.BlackQueensideCastling = false
	}

	// En passant
	if "-" == data[3] {
		pos.EnPassantTargetColumn = -1
		pos.EnPassantTargetRow = -1
	} else {
		pos.EnPassantTargetColumn = charToColumn(data[3][0:1])
		var enPassantRow = charToRow(data[3][1:2])
		if common.Row3 == enPassantRow {
			enPassantRow++
		} else if common.Row6 == enPassantRow {
			enPassantRow--
		}
		pos.EnPassantTargetRow = enPassantRow
	}

	// Clocks
	pos.FiftyMoveClock, _ = strconv.Atoi(data[4])
	pos.FullMoveClock, _ = strconv.Atoi(data[5])

	return pos
}

func charToRow(c string) int {
	var r int
	switch c {
	case "1":
		r = 0
	case "2":
		r = 1
	case "3":
		r = 2
	case "4":
		r = 3
	case "5":
		r = 4
	case "6":
		r = 5
	case "7":
		r = 6
	case "8":
		r = 7
	default:
		panic("Unsupported row: " + c)
	}
	return r
}

func charToColumn(c string) int {
	var r int
	switch c {
	case "a":
		r = 0
	case "b":
		r = 1
	case "c":
		r = 2
	case "d":
		r = 3
	case "e":
		r = 4
	case "f":
		r = 5
	case "g":
		r = 6
	case "h":
		r = 7
	default:
		panic("Unsupported column: " + c)
	}
	return r
}

func charToPiece(c string) int {
	var r int
	switch c {
	case "p":
		r = common.PiecePawn
	case "n":
		r = common.PieceKnight
	case "b":
		r = common.PieceBishop
	case "r":
		r = common.PieceRook
	case "q":
		r = common.PieceQueen
	case "k":
		r = common.PieceKing
	default:
		panic("Unsupported piece to cast: " + c)
	}
	return r
}
