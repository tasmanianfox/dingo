package engine

import (
	"testing"

	"github.com/tasmanianfox/dingo/board"

	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/common"
)

func TestHandleCommand(t *testing.T) {
	var e = new(Engine)
	var c = new(command.QuitCommand)
	var r, s = e.HandleCommand(c)
	if !(true == s && common.ResponseQuit == r.GetType()) {
		t.Errorf("Expected: ResponseQuit")
	}
}

type testPosData struct {
	Position string
	Points   int
	Colour   int
	Depth    int
}

func TestEvaluate(t *testing.T) {
	ps := [...]testPosData{
		testPosData{Position: "rn1qkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", Colour: common.ColourWhite, Depth: 0, Points: MaterialIncrementBishop},
		testPosData{Position: "rn1qkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1", Colour: common.ColourBlack, Depth: 0, Points: MaterialIncrementBishop * -1},
		testPosData{Position: "rn1qkbnr/pppppppp/8/8/4P3/8/PPPPKPPP/RNBQ1BNR w kq - 1 2", Colour: common.ColourWhite, Depth: 0, Points: MaterialIncrementBishop},
		testPosData{Position: "rn1qkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 2", Colour: common.ColourWhite, Depth: 1, Points: MaterialIncrementBishop},
		testPosData{Position: "rn1qkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 2", Colour: common.ColourBlack, Depth: 1, Points: MaterialIncrementBishop * -1},
		testPosData{Position: "rn1qkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 2", Colour: common.ColourWhite, Depth: 2, Points: MaterialIncrementBishop},
		testPosData{Position: "rn1qkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 2", Colour: common.ColourBlack, Depth: 2, Points: MaterialIncrementBishop * -1},
	}

	for i, p := range ps {
		pos := board.FenToPosition(p.Position)
		e := Engine{}

		pts := e.evaluate(pos, p.Colour, p.Depth)
		if p.Points != pts {
			t.Errorf("Expected %d points, but got %d. Pos: %v at %d\n", p.Points, pts, pos, i)
		}
	}

}
