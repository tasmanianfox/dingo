package board

import (
	"testing"

	"github.com/tasmanianfox/dingo/common"
)

const (
	startPositionFen        = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	scandinavianPositionFen = "rnb1kbnr/ppp1pppp/8/3q4/8/8/PPPP1PPP/RNBQKBNR w KQkq - 0 3"
	testCastlingPositionFen = "2kr1bBr/ppp1p1pp/2n5/3q4/6b1/5N2/P2PBPPP/qNBQ1RK1 w - - 0 10"
)

func TestCommitMovement(t *testing.T) {
	// First 2 moves of Scandinavian Defense
	ms := []Movement{
		Movement{SourceColumn: common.ColumnE, SourceRow: common.Row2, DestColumn: common.ColumnE, DestRow: common.Row4},
		Movement{SourceColumn: common.ColumnD, SourceRow: common.Row7, DestColumn: common.ColumnD, DestRow: common.Row5},
		Movement{SourceColumn: common.ColumnE, SourceRow: common.Row4, DestColumn: common.ColumnD, DestRow: common.Row5},
		Movement{SourceColumn: common.ColumnD, SourceRow: common.Row8, DestColumn: common.ColumnD, DestRow: common.Row5},
	}
	testMovements(ms, scandinavianPositionFen, t)

	// Test castlings and promotion
	ms = []Movement{
		Movement{SourceColumn: common.ColumnE, SourceRow: common.Row2, DestColumn: common.ColumnE, DestRow: common.Row4},
		Movement{SourceColumn: common.ColumnD, SourceRow: common.Row7, DestColumn: common.ColumnD, DestRow: common.Row5},
		Movement{SourceColumn: common.ColumnF, SourceRow: common.Row1, DestColumn: common.ColumnE, DestRow: common.Row2},
		Movement{SourceColumn: common.ColumnD, SourceRow: common.Row5, DestColumn: common.ColumnD, DestRow: common.Row4},
		Movement{SourceColumn: common.ColumnC, SourceRow: common.Row2, DestColumn: common.ColumnC, DestRow: common.Row4},
		Movement{SourceColumn: common.ColumnD, SourceRow: common.Row4, DestColumn: common.ColumnC, DestRow: common.Row3}, // en passant
		Movement{SourceColumn: common.ColumnG, SourceRow: common.Row1, DestColumn: common.ColumnF, DestRow: common.Row3},
		Movement{SourceColumn: common.ColumnC, SourceRow: common.Row8, DestColumn: common.ColumnG, DestRow: common.Row4},
		Movement{SourceColumn: common.ColumnE, SourceRow: common.Row1, DestColumn: common.ColumnG, DestRow: common.Row1}, // White kingside castling
		Movement{SourceColumn: common.ColumnB, SourceRow: common.Row8, DestColumn: common.ColumnC, DestRow: common.Row6},
		Movement{SourceColumn: common.ColumnE, SourceRow: common.Row4, DestColumn: common.ColumnE, DestRow: common.Row5},
		Movement{SourceColumn: common.ColumnD, SourceRow: common.Row8, DestColumn: common.ColumnD, DestRow: common.Row5},
		Movement{SourceColumn: common.ColumnE, SourceRow: common.Row5, DestColumn: common.ColumnE, DestRow: common.Row6},
		Movement{SourceColumn: common.ColumnE, SourceRow: common.Row8, DestColumn: common.ColumnC, DestRow: common.Row8}, // Black queenside castling
		Movement{SourceColumn: common.ColumnE, SourceRow: common.Row6, DestColumn: common.ColumnF, DestRow: common.Row7},
		Movement{SourceColumn: common.ColumnC, SourceRow: common.Row3, DestColumn: common.ColumnB, DestRow: common.Row2},
		Movement{SourceColumn: common.ColumnF, SourceRow: common.Row7, DestColumn: common.ColumnG, DestRow: common.Row8, CastTo: common.PieceBishop},
		Movement{SourceColumn: common.ColumnB, SourceRow: common.Row2, DestColumn: common.ColumnA, DestRow: common.Row1, CastTo: common.PieceQueen},
	}
	testMovements(ms, testCastlingPositionFen, t)
}

func testMovements(ms []Movement, fen string, t *testing.T) {
	var p = FenToPosition(startPositionFen)
	var ps = PositionToFen(p)
	if ps != startPositionFen {
		t.Errorf("Invalid convertion of FEN position. Expected: '%s', got: '%s'", startPositionFen, ps)
	}

	for _, m := range ms {
		p = CommitMovement(p, m)
	}
	ps = PositionToFen(p)
	if ps != fen {
		t.Errorf("Invalid convertion of FEN position. Expected: '%s', got: '%s'", fen, ps)
	}
}
