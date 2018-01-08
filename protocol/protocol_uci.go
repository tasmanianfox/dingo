package protocol

import (
	"bufio"
	"strings"

	"github.com/tasmanianfox/dingo/board"
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
	p.Output(new(response.AppNameResponse))
	p.Output(new(response.AuthorResponse))
	return p
}

func (p UciProtocol) ReadCommand() (command.Command, bool) {
	var c command.Command
	var s = false
	p.BufferedScanner.Scan()
	var i = p.BufferedScanner.Text()
	var ia = strings.Split(i, " ")
	if len(ia) > 0 {
		s = true
		var sc = ia[0]
		var arg = ia[1:]
		switch sc {
		case "isready":
			c = new(command.CheckIfIsEngineReadyCommand)
		case "position":
			c = p.getSetPositionCommand(arg)
		case "quit":
			c = new(command.QuitCommand)
		default:
			s = false
		}
	}

	return c, s
}

func (p UciProtocol) Output(r response.Response) {
	var f = true
	var s = ""

	switch r.GetType() {
	case common.ResponseReady:
		s = "readyok"
	case common.ResponseAppName:
		s = "id name DinGo"
	case common.ResponseAuthor:
		s = "id author Sergei Belyakov"
	default:
		f = false
	}

	if true == f {
		p.BufferedWriter.WriteString(s + "\n")
		p.BufferedWriter.Flush()
	}
}

func (p UciProtocol) getSetPositionCommand(args []string) command.Command {
	var c command.SetPositionCommand
	for i, arg := range args {
		if "startpos" == arg {
			c.NewGame = true
		} else if 1 == i && "moves" == arg {
			c.Mode = command.SetPositionModeMoves
		} else {
			var m = p.stringToMovement(arg)
			c.AddMovement(m)
		}
	}

	return c
}

func (p UciProtocol) stringToMovement(s string) board.Movement {
	if len(s) < 4 {
		panic("Movement must be at least 4 characters long")
	}
	var m board.Movement
	m.SourceVertical = p.charToVertical(s[0:1])
	m.SourceHorizontal = p.charToHorizontal(s[1:2])
	m.DestVertical = p.charToVertical(s[2:3])
	m.DestHorizontal = p.charToHorizontal(s[3:4])
	if 5 == len(s) {
		m.CastTo = p.charToPieceToCast(s[4:5])
	}
	return m
}

func (p UciProtocol) charToHorizontal(c string) int {
	var r int
	switch c {
	case "1":
		r = 0
	case "2":
		r = 1
	case "3":
		r = 2
	case "4":
		r = 3
	case "5":
		r = 4
	case "6":
		r = 5
	case "7":
		r = 6
	case "8":
		r = 7
	default:
		panic("Unsupported horizontal: " + c)
	}
	return r
}

func (p UciProtocol) charToVertical(c string) int {
	var r int
	switch c {
	case "a":
		r = 0
	case "b":
		r = 1
	case "c":
		r = 2
	case "d":
		r = 3
	case "e":
		r = 4
	case "f":
		r = 5
	case "g":
		r = 6
	case "h":
		r = 7
	default:
		panic("Unsupported vertical: " + c)
	}
	return r
}

func (p UciProtocol) charToPieceToCast(c string) int {
	var r int
	switch c {
	case "b":
		r = common.PieceBishop
	case "k":
		r = common.PieceKnight
	case "r":
		r = common.PieceRook
	case "q":
		r = common.PieceQueen
	default:
		panic("Unsupported piece to cast: " + c)
	}
	return r
}
