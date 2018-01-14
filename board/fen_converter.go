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

// PositionToFen converts an instance of Position struct to FEN string representation
func PositionToFen(p Position) string {
	f := ""
	// Pieces
	for i, r := range p.Board {
		e := 0   // num of empty cells
		fb := "" // buffer for rows
		for _, pc := range r {
			if common.ColourEmpty == pc.Colour && common.PieceEmpty == pc.Type {
				e++
			} else {
				if e > 0 { // add the number of empty cells detected previously
					fb += strconv.Itoa(e)
					e = 0
				}
				fb += pieceToChar(pc)
			}
		}
		if e > 0 { // add the number of empty cells at the end of row if exists
			fb += strconv.Itoa(e)
		}
		if i < len(p.Board)-1 { // add a delimiter between riws
			fb = "/" + fb
		}
		f = fb + f
	}
	f += " "
	if common.ColourWhite == p.ActiveColour {
		f += "w "
	} else if common.ColourBlack == p.ActiveColour {
		f += "b "
	} else {
		panic("Unsupported color: " + strconv.Itoa(p.ActiveColour))
	}
	// Castling
	c := ""
	if p.WhiteKingsideCastling {
		c += "K"
	}
	if p.WhiteQueensideCastling {
		c += "Q"
	}
	if p.BlackKingsideCastling {
		c += "k"
	}
	if p.BlackQueensideCastling {
		c += "q"
	}
	if 0 == len(c) {
		c = "-"
	}
	f += c + " "

	// En passant
	if p.EnPassantTargetColumn == -1 && p.EnPassantTargetRow == -1 {
		f += "-"
	} else {
		var row = -1
		if common.Row4 == p.EnPassantTargetRow {
			row = common.Row3
		} else if common.Row5 == p.EnPassantTargetRow {
			row = common.Row6
		} else {
			panic("Invalid row for en passant: " + strconv.Itoa(p.EnPassantTargetRow))
		}
		f += columnToChar(p.EnPassantTargetColumn) + rowToChar(row)
	}
	f += " " + strconv.Itoa(p.FiftyMoveClock) + " " + strconv.Itoa(p.FullMoveClock)

	return f
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

func rowToChar(r int) string {
	var c string
	switch r {
	case 0:
		c = "1"
	case 1:
		c = "2"
	case 2:
		c = "3"
	case 3:
		c = "4"
	case 4:
		c = "5"
	case 5:
		c = "6"
	case 6:
		c = "7"
	case 7:
		c = "8"
	default:
		panic("Unsupported row: " + strconv.Itoa(r))
	}
	return c
}

func columnToChar(r int) string {
	var c string
	switch r {
	case 0:
		c = "a"
	case 1:
		c = "b"
	case 2:
		c = "c"
	case 3:
		c = "d"
	case 4:
		c = "e"
	case 5:
		c = "f"
	case 6:
		c = "g"
	case 7:
		c = "h"
	default:
		panic("Unsupported column: " + strconv.Itoa(r))
	}
	return c
}

func pieceToChar(p Piece) string {
	var s string
	switch p.Type {
	case common.PiecePawn:
		s = "p"
	case common.PieceKnight:
		s = "n"
	case common.PieceBishop:
		s = "b"
	case common.PieceRook:
		s = "r"
	case common.PieceQueen:
		s = "q"
	case common.PieceKing:
		s = "k"
	default:
		panic("Unsupported piece to cast: (" + strconv.Itoa(p.Colour) + "/" + strconv.Itoa(p.Type) + ")")
	}
	if common.ColourWhite == p.Colour {
		s = strings.ToUpper(s)
	}
	return s
}
