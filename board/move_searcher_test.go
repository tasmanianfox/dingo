package board

import (
	"testing"

	"github.com/tasmanianfox/dingo/common"
)

type kingPosition struct {
	Fen    string
	Column int
	Row    int
	Moves  int
}

func TestFindKingMoves(t *testing.T) {
	pds := [...]kingPosition{
		kingPosition{Fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", Column: common.ColumnE, Row: common.Row1, Moves: 0},
		kingPosition{Fen: "r2qk2r/pbp2ppp/1pnp1n2/1Bb1p1B1/4P3/2NP1N2/PPPQ1PPP/R3K2R w KQkq - 0 8", Column: common.ColumnE, Row: common.Row1, Moves: 5},
		kingPosition{Fen: "r2qk2r/pbp2ppp/1pnp1n2/1Bb1p1B1/4P3/2NP1N1P/PPPQ1PP1/R3K2R b KQkq - 0 8", Column: common.ColumnE, Row: common.Row8, Moves: 4},
		kingPosition{Fen: "r3k2r/p1pqppbp/1pnp1np1/8/6b1/2NPPNPP/PPPBQPB1/R3K2R b K - 1 11", Column: common.ColumnE, Row: common.Row8, Moves: 2},
		kingPosition{Fen: "r3k2r/p1pqppb1/1pnp1npp/8/6b1/2NPPNPP/PPPBQPB1/R3K2R w K - 0 12", Column: common.ColumnE, Row: common.Row1, Moves: 3},
		kingPosition{Fen: "4r3/b4N2/8/2P4k/8/5K2/8/8 w - - 0 1", Column: common.ColumnF, Row: common.Row3, Moves: 4},
		kingPosition{Fen: "4r3/b4N2/8/2P4k/8/5K2/8/8 b - - 0 1", Column: common.ColumnH, Row: common.Row5, Moves: 2},
		kingPosition{Fen: "r3k2r/8/N7/1b6/8/8/8/R3K2R w KQkq - 0 1", Column: common.ColumnE, Row: common.Row1, Moves: 4},
		kingPosition{Fen: "r3k2r/8/N7/1b6/8/8/8/R3K2R b KQkq - 0 1", Column: common.ColumnE, Row: common.Row8, Moves: 7},
		kingPosition{Fen: "rnbqk2r/pppp1ppp/4Q2n/2b1p3/4P3/8/PPPPBPPP/RNB1K1NR b KQkq - 0 1", Column: common.ColumnE, Row: common.Row8, Moves: 1},
	}

	for i, pd := range pds {
		p := FenToPosition(pd.Fen)
		ms := findKingMoves(p, pd.Row, pd.Column)
		if len(ms) != pd.Moves {
			t.Errorf("Incorrect number of available moves (%d)", i)
		}
	}
}
