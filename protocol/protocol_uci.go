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

func (p UciProtocol) getSetPositionCommand(args []string) command.SetPositionCommand {
	var c command.SetPositionCommand
	if len(args) >= 3 && "startpos" == args[0] && "moves" == args[1] {
		c = p.getSetPositionCommandFromMovements(args[2:])
	} else if 7 == len(args) && "fen" == args[0] {
		c = p.getSetPositionCommandFromFEN(strings.Join(args[1:], " "))
	}

	return c
}

func (p UciProtocol) getSetPositionCommandFromFEN(s string) command.SetPositionCommand {
	c := command.SetPositionCommand{NewGame: false, Mode: command.SetPositionModeDirect}
	c.Position = board.FenToPosition(s)
	return c
}

func (p UciProtocol) getSetPositionCommandFromMovements(movements []string) command.SetPositionCommand {
	c := command.SetPositionCommand{NewGame: true, Mode: command.SetPositionModeMovements}
	for _, ms := range movements {
		var m = p.stringToMovement(ms)
		c.AddMovement(m)
	}
	return c
}

// Replacements

func (p UciProtocol) stringToMovement(s string) board.Movement {
	if len(s) < 4 {
		panic("Movement must be at least 4 characters long")
	}
	var m board.Movement
	m.SourceColumn = p.charToColumn(s[0:1])
	m.SourceRow = p.charToRow(s[1:2])
	m.DestColumn = p.charToColumn(s[2:3])
	m.DestRow = p.charToRow(s[3:4])
	if 5 == len(s) {
		m.CastTo = p.charToPiece(s[4:5])
	}
	return m
}

func (p UciProtocol) charToRow(c string) int {
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
		panic("Unsupported row: " + c)
	}
	return r
}

func (p UciProtocol) charToColumn(c string) int {
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
		panic("Unsupported column: " + c)
	}
	return r
}

func (p UciProtocol) charToPiece(c string) int {
	var r int
	switch c {
	case "p":
		r = common.PiecePawn
	case "n":
		r = common.PieceKnight
	case "b":
		r = common.PieceBishop
	case "r":
		r = common.PieceRook
	case "q":
		r = common.PieceQueen
	case "k":
		r = common.PieceKing
	default:
		panic("Unsupported piece to cast: " + c)
	}
	return r
}
