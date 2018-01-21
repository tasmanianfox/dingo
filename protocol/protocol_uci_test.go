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
