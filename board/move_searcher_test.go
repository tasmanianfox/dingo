package board

import (
	"testing"

	"github.com/tasmanianfox/dingo/common"
)

type testData struct {
	Fen    string
	Column int
	Row    int
	Moves  int
}

func TestFindKingMoves(t *testing.T) {
	pds := [...]testData{
		testData{Fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", Column: common.ColumnE, Row: common.Row1, Moves: 0},
		testData{Fen: "r2qk2r/pbp2ppp/1pnp1n2/1Bb1p1B1/4P3/2NP1N2/PPPQ1PPP/R3K2R w KQkq - 0 8", Column: common.ColumnE, Row: common.Row1, Moves: 5},
		testData{Fen: "r2qk2r/pbp2ppp/1pnp1n2/1Bb1p1B1/4P3/2NP1N1P/PPPQ1PP1/R3K2R b KQkq - 0 8", Column: common.ColumnE, Row: common.Row8, Moves: 4},
		testData{Fen: "r3k2r/p1pqppbp/1pnp1np1/8/6b1/2NPPNPP/PPPBQPB1/R3K2R b K - 1 11", Column: common.ColumnE, Row: common.Row8, Moves: 2},
		testData{Fen: "r3k2r/p1pqppb1/1pnp1npp/8/6b1/2NPPNPP/PPPBQPB1/R3K2R w K - 0 12", Column: common.ColumnE, Row: common.Row1, Moves: 3},
		testData{Fen: "4r3/b4N2/8/2P4k/8/5K2/8/8 w - - 0 1", Column: common.ColumnF, Row: common.Row3, Moves: 4},
		testData{Fen: "4r3/b4N2/8/2P4k/8/5K2/8/8 b - - 0 1", Column: common.ColumnH, Row: common.Row5, Moves: 2},
		testData{Fen: "r3k2r/8/N7/1b6/8/8/8/R3K2R w KQkq - 0 1", Column: common.ColumnE, Row: common.Row1, Moves: 4},
		testData{Fen: "r3k2r/8/N7/1b6/8/8/8/R3K2R b KQkq - 0 1", Column: common.ColumnE, Row: common.Row8, Moves: 7},
		testData{Fen: "rnbqk2r/pppp1ppp/4Q2n/2b1p3/4P3/8/PPPPBPPP/RNB1K1NR b KQkq - 0 1", Column: common.ColumnE, Row: common.Row8, Moves: 1},
		testData{Fen: "8/8/5P2/4Kb2/3r4/4n3/8/4k3 w - - 0 1", Column: common.ColumnE, Row: common.Row5, Moves: 1},
	}

	for i, pd := range pds {
		p := FenToPosition(pd.Fen)
		ms := findKingMoves(p, pd.Row, pd.Column)
		if len(ms) != pd.Moves {
			t.Errorf("Incorrect number of available moves (%d / %d) %x", i, len(ms), ms)
		}
	}
}

func TestFindUsualPieceMoves(t *testing.T) {
	pds := [...]testData{
		testData{Fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", Column: common.ColumnG, Row: common.Row1, Moves: 2},
		testData{Fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", Column: common.ColumnA, Row: common.Row1, Moves: 0},
		testData{Fen: "r2qk1nr/p2p1pp1/bpn1p2p/6B1/1b1NP3/2N5/PPP1BPPP/R2Q1RK1 w kq - 0 9", Column: common.ColumnE, Row: common.Row2, Moves: 7},
		testData{Fen: "r2qk1nr/p2p1pp1/bpn1p2p/8/1b1NP2B/2N5/PPP1BPPP/R2Q1RK1 b kq - 1 9", Column: common.ColumnC, Row: common.Row6, Moves: 5},
		testData{Fen: "8/5k2/8/7K/2q5/1B6/8/8 b - - 0 1", Column: common.ColumnC, Row: common.Row4, Moves: 3},
	}

	for i, pd := range pds {
		p := FenToPosition(pd.Fen)
		ms := findUsualPieceMoves(p, pd.Row, pd.Column)
		if len(ms) != pd.Moves {
			t.Errorf("Incorrect number of available moves (%d) %d %x", i, len(ms), ms)
		}
	}
}

func TestFindPawnPieceMoves(t *testing.T) {
	pds := [...]testData{
		testData{Fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", Column: common.ColumnG, Row: common.Row2, Moves: 2},
		testData{Fen: "rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1", Column: common.ColumnE, Row: common.Row4, Moves: 2},
		testData{Fen: "rnbqkbnr/ppp1pppp/8/4P3/2Pp4/8/PP1P1PPP/RNBQKBNR b KQkq c3 0 1", Column: common.ColumnD, Row: common.Row4, Moves: 2},
		testData{Fen: "rnbqkbnr/pp2p1pp/4P3/3p1p2/5P2/2p5/PPPP2PP/RNBQKBNR w KQkq - 0 1", Column: common.ColumnE, Row: common.Row6, Moves: 0},
		testData{Fen: "rnbqkbnr/ppp1p1pp/8/3pPp2/8/8/PPPP1PPP/RNBQKBNR w KQkq f6 0 1", Column: common.ColumnE, Row: common.Row5, Moves: 2},
		testData{Fen: "8/5k1K/8/8/8/8/5p2/8 b - - 0 1", Column: common.ColumnF, Row: common.Row2, Moves: 4},
	}

	for i, pd := range pds {
		p := FenToPosition(pd.Fen)
		ms := findPawnMoves(p, pd.Row, pd.Column)
		if len(ms) != pd.Moves {
			t.Errorf("Incorrect number of available moves (%d) %d %x", i, len(ms), ms)
		}
	}
}
