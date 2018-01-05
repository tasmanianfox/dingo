package protocol

import (
	"bufio"

	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/common"
	"github.com/tasmanianfox/dingo/response"
)

type UciProtocol struct {
	BufferedScanner *bufio.Scanner
	BufferedWriter  *bufio.Writer
}

func NewUciProtocol(s *bufio.Scanner, w *bufio.Writer) UciProtocol {
	var p = UciProtocol{BufferedScanner: s, BufferedWriter: w}
	return p
}

func (p UciProtocol) ReadCommand() (command.Command, bool) {
	var c command.Command
	var s bool = true
	p.BufferedScanner.Scan()
	var i string = p.BufferedScanner.Text()
	switch i {
	case "isready":
		c = new(command.CheckIfIsEngineReadyCommand)
	case "quit":
		c = new(command.QuitCommand)
	default:
		s = false
	}

	return c, s
}

func (p UciProtocol) Output(r response.Response) {
	var f bool = true
	var s string = ""

	switch r.GetType() {
	case common.ResponseReady:
		s = "readyok"
	default:
		f = false
	}

	if true == f {
		p.BufferedWriter.WriteString(s + "\n")
		p.BufferedWriter.Flush()
	}
}
