package protocol

import (
	"bufio"

	"github.com/tasmanianfox/dingo/command"
)

type UciProtocol struct {
	BufferedScanner *bufio.Scanner
}

func NewUciProtocol(s *bufio.Scanner) UciProtocol {
	var p = UciProtocol{BufferedScanner: s}
	return p
}

func (p UciProtocol) ReadCommand() command.Command {
	var c command.Command
	p.BufferedScanner.Scan()
	var i string = p.BufferedScanner.Text()
	switch i {
	case "isready":
		c = new(command.CheckIfIsEngineReadyCommand)
	case "quit":
		c = new(command.QuitCommand)
	}

	return c
}
