package protocol

import (
	"bufio"
	"io"
	"os"
	"strconv"
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
	if !c2.NewGame {
		t.Errorf("The NewGame flag is false, expected: true")
	}
	if len(c2.Movements) != 3 {
		t.Errorf("Expected number of movements: 3, got: " + strconv.Itoa(len(c2.Movements)))
	}
	var movements = [3][4]int{}
	movements[0][0] = common.ColumnE
	movements[0][1] = common.Row2
	movements[0][2] = common.ColumnE
	movements[0][3] = common.Row4
	movements[1][0] = common.ColumnD
	movements[1][1] = common.Row7
	movements[1][2] = common.ColumnD
	movements[1][3] = common.Row5
	movements[2][0] = common.ColumnE
	movements[2][1] = common.Row4
	movements[2][2] = common.ColumnD
	movements[2][3] = common.Row5
	for i, m := range c2.Movements {
		if !(m.SourceColumn == movements[i][0] && m.SourceRow == movements[i][1]) {
			t.Errorf("Incorrect source cell: (" + strconv.Itoa(movements[i][0]) + ", " + strconv.Itoa(movements[i][1]) + ") expected" +
				", but got: (" + strconv.Itoa(m.SourceColumn) + ", " + strconv.Itoa(m.SourceRow) + "). Movement: " + strconv.Itoa(i))
		}
		if !(m.DestColumn == movements[i][2] && m.DestRow == movements[i][3]) {
			t.Errorf("Incorrect dest cell: (" + strconv.Itoa(movements[i][2]) + ", " + strconv.Itoa(movements[i][3]) + ") expected" +
				", but got: (" + strconv.Itoa(m.DestColumn) + ", " + strconv.Itoa(m.DestRow) + "). Movement: " + strconv.Itoa(i))
		}
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
