package board

import (
	"testing"

	"github.com/tasmanianfox/dingo/common"
)

const (
	startPositionFen        = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	scandinavianPositionFen = "rnb1kbnr/ppp1pppp/8/3q4/8/8/PPPP1PPP/RNBQKBNR w KQkq - 0 3"
)

func TestCommitMovement(t *testing.T) {
	// First 2 moves of Scandinavian Defense
	var p = FenToPosition(startPositionFen)
	var ps = PositionToFen(p)
	if ps != startPositionFen {
		t.Errorf("Invalid convertion of FEN position. Expected: '%s', got: '%s'", startPositionFen, ps)
	}
	ms := [...]Movement{
		Movement{SourceColumn: common.ColumnE, SourceRow: common.Row2, DestColumn: common.ColumnE, DestRow: common.Row4},
		Movement{SourceColumn: common.ColumnD, SourceRow: common.Row7, DestColumn: common.ColumnD, DestRow: common.Row5},
		Movement{SourceColumn: common.ColumnE, SourceRow: common.Row4, DestColumn: common.ColumnD, DestRow: common.Row5},
		Movement{SourceColumn: common.ColumnD, SourceRow: common.Row8, DestColumn: common.ColumnD, DestRow: common.Row5},
	}
	for _, m := range ms {
		p = CommitMovement(p, m)
	}
	ps = PositionToFen(p)
	if ps != scandinavianPositionFen {
		t.Errorf("Invalid convertion of FEN position. Expected: '%s', got: '%s'", scandinavianPositionFen, ps)
	}
}
