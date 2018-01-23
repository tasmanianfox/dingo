package protocol

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/tasmanianfox/dingo/board"

	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/common"
)

func TestReadCommand(t *testing.T) {
	var r io.Reader = strings.NewReader("isready\n")
	var s = bufio.NewScanner(r)
	var w = bufio.NewWriter(os.Stdout)
	var p = NewUciProtocol(s, w)

	var c, success = p.ReadCommand()
	if !(true == success && common.СommandCheckIfIsEngineReady == c.GetType()) {
		t.Errorf("Expected: CheckIfIsEngineReady command")
	}

	r = strings.NewReader("quit\n")
	s = bufio.NewScanner(r)
	p = NewUciProtocol(s, w)
	c, success = p.ReadCommand()
	if !(true == success && common.СommandQuit == c.GetType()) {
		t.Errorf("Expected: Quit command")
	}

	r = strings.NewReader("unknown_command\n")
	s = bufio.NewScanner(r)
	p = NewUciProtocol(s, w)
	c, success = p.ReadCommand()
	if true == success {
		t.Errorf("Expected: nil")
	}

	r = strings.NewReader("position startpos moves e2e4 d7d5 e4d5\n")
	s = bufio.NewScanner(r)
	p = NewUciProtocol(s, w)
	c, success = p.ReadCommand()
	if !(true == success && common.СommandSetPosition == c.GetType()) {
		t.Errorf("Expected: Quit command")
	}
	c2, ok := c.(command.SetPositionCommand)
	if !ok {
		t.Errorf("Expected type: command.SetPositionCommand")
	}
	if common.ColourBlack != c2.Position.ActiveColour {
		t.Errorf("Expected active colour: black")
	}

	r = strings.NewReader("position fen rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - e3 0 1\n")
	s = bufio.NewScanner(r)
	p = NewUciProtocol(s, w)
	c, success = p.ReadCommand()
	c2, ok = c.(command.SetPositionCommand)

	if !(!c2.Position.WhiteKingsideCastling && !c2.Position.WhiteQueensideCastling && !c2.Position.BlackKingsideCastling && !c2.Position.BlackQueensideCastling) {
		t.Errorf("Expected: none castlings available. Actual: WK %t WQ %t BK %t BQ %t", c2.Position.WhiteKingsideCastling, c2.Position.WhiteQueensideCastling,
			c2.Position.BlackKingsideCastling, c2.Position.BlackQueensideCastling)
	}

	positions := [...][2]string{
		{
			"position startpos moves e2e4 f7f6 g2g3 e8f7 f1c4\n",
			"rnbq1bnr/pppppkpp/5p2/8/2B1P3/6P1/PPPP1P1P/RNBQK1NR b KQ - 2 3",
		},
		{
			"position startpos moves e2e4 b8c6 h2h4 c6b4 e1e2 b4c2 a2a3 c2a3 d2d4 e7e6 b1d2 a3b1 a1a6 g7g6 h1h3 g8h6 d1a4 d8g5 a6a5 c7c5 a4c6 f8g7",
			"r1b1k2r/pp1p1pbp/2Q1p1pn/R1p3q1/3PP2P/7R/1P1NKPP1/1nB2BN1 w kq - 2 12",
		},
		{
			"position startpos moves g2g4 b8a6 c2c4 c7c5 d1c2 e7e6 c2c3 f7f5 c3e5 f5f4 f1g2 g8h6 g2f1 f4f3 e5c7 d7d5 c7g3 e8e7 g3e5 d8d6 e5c3 e6e5 h2h4 b7b5 d2d4 d6f6 f1g2 e5d4 c3c2 f6h4 c1h6 h4h3 c2c1 h3g3 c1c3 a6c7 e2f3 c7e6 g2f1 e7f7 h1h3 g3d6 c3e3 d6a6 e3e5 a6a4 f1e2 a4c2 e5e6 f7e6 a2a4 c2d3 c4b5 d3b5 b1d2 c5c4 e1c1 g7g6 b2b3 c8b7 e2d3 f8b4 d3g6 b5b6 c1c2 b4c3 h3g3 b6c6 b3c4 c6d6 g3h3 d6f8 c2b3 e6e7 g6f7 f8c8 d1e1",
			"r1q4r/pb2kB1p/7B/3p4/P1Pp2P1/1Kb2P1R/3N1P2/4R1N1 b - - 8 39",
		},
	}
	for i, pData := range positions {
		pos := pData[0]
		fen := pData[1]

		r = strings.NewReader(pos)
		s = bufio.NewScanner(r)
		p = NewUciProtocol(s, w)
		c, success = p.ReadCommand()
		c2, ok = c.(command.SetPositionCommand)
		apFen := board.PositionToFen(c2.Position)
		if apFen != fen {
			t.Errorf("Expected (%d): \n%s\n, got: \n%s\n", i, fen, apFen)
		}
	}
}
