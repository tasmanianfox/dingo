package protocol

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

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

	r = strings.NewReader("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b - e3 0 1\n")
	s = bufio.NewScanner(r)
	p = NewUciProtocol(s, w)
	c, success = p.ReadCommand()
	c2, ok = c.(command.SetPositionCommand)

	if !(!c2.Position.WhiteKingsideCastling && !c2.Position.WhiteQueensideCastling && !c2.Position.BlackKingsideCastling && !c2.Position.BlackQueensideCastling) {
		t.Errorf("Expected: none castlings available. Actual: WK %t WQ %t BK %t BQ %t", c2.Position.WhiteKingsideCastling, c2.Position.WhiteQueensideCastling,
			c2.Position.BlackKingsideCastling, c2.Position.BlackQueensideCastling)
	}
}
