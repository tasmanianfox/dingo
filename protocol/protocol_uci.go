package protocol

import (
	"bufio"
	"fmt"
	"strconv"
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
	p.Output(new(response.ProtocolConfirmationResponse))
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
		case "go":
			c = command.CalculateMoveCommand{}
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
	case common.ResponseProtocolConfirmation:
		s = "uciok"
	case common.ResponseMove:
		r2 := r.(response.MoveResponse)
		m := r2.Move
		cp := ""
		if m.CastTo > -1 {
			cp = p.pieceToChar(m.CastTo)
		}
		s = fmt.Sprintf(
			"bestmove %s%s%s%s%s",
			p.columnToChar(m.SourceColumn),
			p.rowToChar(m.SourceRow),
			p.columnToChar(m.DestColumn),
			p.rowToChar(m.DestRow),
			cp,
		)

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
		c = p.getSetPositionCommandFromMoves(args[2:])
	} else if 7 == len(args) && "fen" == args[0] {
		c = p.getSetPositionCommandFromFEN(strings.Join(args[1:], " "))
	}

	return c
}

func (p UciProtocol) getSetPositionCommandFromFEN(s string) command.SetPositionCommand {
	var c command.SetPositionCommand
	c.Position = board.FenToPosition(s)
	return c
}

func (p UciProtocol) getSetPositionCommandFromMoves(moves []string) command.SetPositionCommand {
	var c command.SetPositionCommand
	c.Position = board.BuildStartPosition()
	for _, ms := range moves {
		var m = p.stringToMove(ms)
		c.Position = board.CommitMove(c.Position, m)
	}
	return c
}

// Replacements

func (p UciProtocol) stringToMove(s string) board.Move {
	if len(s) < 4 {
		panic("Move must be at least 4 characters long")
	}
	var m board.Move
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
		r = common.Row1
	case "2":
		r = common.Row2
	case "3":
		r = common.Row3
	case "4":
		r = common.Row4
	case "5":
		r = common.Row5
	case "6":
		r = common.Row6
	case "7":
		r = common.Row7
	case "8":
		r = common.Row8
	default:
		panic("Unsupported row: " + c)
	}
	return r
}

func (p UciProtocol) charToColumn(c string) int {
	var r int
	switch c {
	case "a":
		r = common.ColumnA
	case "b":
		r = common.ColumnB
	case "c":
		r = common.ColumnC
	case "d":
		r = common.ColumnD
	case "e":
		r = common.ColumnE
	case "f":
		r = common.ColumnF
	case "g":
		r = common.ColumnG
	case "h":
		r = common.ColumnH
	default:
		panic("Unsupported column: " + c)
	}
	return r
}

func (p UciProtocol) rowToChar(r int) string {
	var res string
	switch r {
	case common.Row1:
		res = "1"
	case common.Row2:
		res = "2"
	case common.Row3:
		res = "3"
	case common.Row4:
		res = "4"
	case common.Row5:
		res = "5"
	case common.Row6:
		res = "6"
	case common.Row7:
		res = "7"
	case common.Row8:
		res = "8"
	default:
		panic("Unsupported row: " + strconv.Itoa(r))
	}
	return res
}

func (p UciProtocol) columnToChar(c int) string {
	var res string
	switch c {
	case common.ColumnA:
		res = "a"
	case common.ColumnB:
		res = "b"
	case common.ColumnC:
		res = "c"
	case common.ColumnD:
		res = "d"
	case common.ColumnE:
		res = "e"
	case common.ColumnF:
		res = "f"
	case common.ColumnG:
		res = "g"
	case common.ColumnH:
		res = "h"
	default:
		panic("Unsupported column: " + strconv.Itoa(c))
	}
	return res
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

func (p UciProtocol) pieceToChar(c int) string {
	var res string
	switch c {
	case common.PiecePawn:
		res = "p"
	case common.PieceKnight:
		res = "n"
	case common.PieceBishop:
		res = "b"
	case common.PieceRook:
		res = "r"
	case common.PieceQueen:
		res = "q"
	case common.PieceKing:
		res = "k"
	default:
		panic("Unsupported piece to cast: " + strconv.Itoa(c))
	}
	return res
}
