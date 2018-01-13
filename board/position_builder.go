package board

import "github.com/tasmanianfox/dingo/common"

// BuildEmptyPosition initializes an empty position with no pieces
func BuildEmptyPosition() Position {
	p := Position{ActiveColour: common.ColourEmpty,
		WhiteKingsideCastling: false, WhiteQueensideCastling: false,
		BlackKingsideCastling: false, BlackQueensideCastling: false,
		EnPassantTargetColumn: -1, EnPassantTargetRow: -1,
		FiftyMoveClock: 0, FullMoveClock: 0}
	var pc = Piece{Colour: common.ColourEmpty, Type: common.PieceEmpty}
	for i, r := range p.Board {
		for j := range r {
			p.Board[i][j] = pc
		}
	}
	return p
}

// BuildStartPosition initializes a start position
func BuildStartPosition() Position {
	p := BuildEmptyPosition()
	pts := [...]int{common.PieceRook, common.PieceKnight, common.PieceBishop, common.PieceQueen,
		common.PieceKing, common.PieceBishop, common.PieceKnight, common.PieceRook}
	for i, pt := range pts {
		p.Board[common.Row1][i] = Piece{Colour: common.ColourWhite, Type: pt}
		p.Board[common.Row2][i] = Piece{Colour: common.ColourWhite, Type: common.PiecePawn}
		p.Board[common.Row8][i] = Piece{Colour: common.ColourBlack, Type: pt}
		p.Board[common.Row2][i] = Piece{Colour: common.ColourBlack, Type: common.PiecePawn}
	}

	return p
}
