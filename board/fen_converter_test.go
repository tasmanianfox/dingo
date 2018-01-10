package board

import (
	"testing"

	"github.com/tasmanianfox/dingo/common"
)

func TestRun(t *testing.T) {
	var p = FenToPosition("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	if p.ActiveColour != common.ColourWhite {
		t.Errorf("Expected colour: white, got: %d", p.ActiveColour)
	}
	if !(p.WhiteKingsideCastling && p.WhiteQueensideCastling && p.BlackKingsideCastling && p.BlackQueensideCastling) {
		t.Errorf("Expected: all castlings available. Actual: WK %t WQ %t BK %t BQ %t", p.WhiteKingsideCastling, p.WhiteQueensideCastling,
			p.BlackKingsideCastling, p.BlackQueensideCastling)
	}
	if !(common.ColourWhite == p.Board[0][6].Colour && common.PieceKnight == p.Board[0][6].Type) {
		t.Errorf("Expected a white knight on G1, got: %d %d", p.Board[0][6].Colour, p.Board[0][6].Type)
	}
	if !(common.ColourBlack == p.Board[7][2].Colour && common.PieceBishop == p.Board[7][2].Type) {
		t.Errorf("Expected a black bishop on C8, got: %d %d", p.Board[7][2].Colour, p.Board[7][2].Type)
	}
	if !(common.ColourEmpty == p.Board[4][3].Colour && common.PieceEmpty == p.Board[4][3].Type) {
		t.Errorf("Expected an empty cell on D5, got: %d %d", p.Board[4][3].Colour, p.Board[4][3].Type)
	}

	p = FenToPosition("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - e3 0 1\n")
	if p.ActiveColour != common.ColourBlack {
		t.Errorf("Expected colour: black, got: %d", p.ActiveColour)
	}
	if !(!p.WhiteKingsideCastling && !p.WhiteQueensideCastling && !p.BlackKingsideCastling && !p.BlackQueensideCastling) {
		t.Errorf("Expected: none castlings available. Actual: WK %t WQ %t BK %t BQ %t", p.WhiteKingsideCastling, p.WhiteQueensideCastling,
			p.BlackKingsideCastling, p.BlackQueensideCastling)
	}
}
