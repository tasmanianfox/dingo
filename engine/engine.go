package engine

import (
	"math/rand"
	"time"

	"github.com/tasmanianfox/dingo/board"
	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/common"
	"github.com/tasmanianfox/dingo/response"
)

const (
	MaterialIncrementPawn   = 100
	MaterialIncrementKnight = 300
	MaterialIncrementBishop = 300
	MaterialIncrementRook   = 500
	MaterialIncrementQueen  = 900
)

type Engine struct {
	Position        board.Position
	ResponseChannel chan response.Response
}

// HandleCommand Handles given command and returns an instance of Response object
func (e *Engine) HandleCommand(c command.Command) (response.Response, bool) {
	var r response.Response
	var s = true
	switch c.GetType() {
	case common.СommandQuit:
		r = new(response.QuitResponse)
	case common.СommandCheckIfIsEngineReady:
		r = new(response.ReadyResponse)
	case common.СommandSetPosition:
		c2 := c.(command.SetPositionCommand)
		e.Position = c2.Position
		r = new(response.EmptyResponse)
	case common.СommandCalculateMove:
		c2 := c.(command.CalculateMoveCommand)
		go e.calculateMove(c2)
		r = new(response.EmptyResponse)
	default:
		s = false
	}

	return r, s
}

func (e Engine) calculateMove(c command.CalculateMoveCommand) {
	ms := board.FindAllAvailableMoves(e.Position)
	bms := []board.Move{}
	bs := common.MinusInf

	p := e.Position
	for _, m := range ms {
		p2 := board.CommitMove(p, m)
		s := e.evaluate(p2, p.ActiveColour, 3)
		if s > bs {
			bms = []board.Move{}
			bs = s
		}
		if s >= bs {
			bms = append(bms, m)
		}
	}

	rand.Seed(time.Now().UnixNano())
	numMs := len(bms)
	m := board.Move{}
	if numMs > 1 {
		m = bms[rand.Intn(len(bms)-1)]
	} else if numMs == 1 {
		m = bms[0]
	} else {
		panic("Cannot find any moves")
	}
	r := response.MoveResponse{Move: m}
	e.ResponseChannel <- r
}

func (e Engine) evaluate(pos board.Position, colour int, depth int) int {
	r := common.MinusInf

	points := 0
	if pos.IsKingCheckmated(pos.ActiveColour) {
		if pos.ActiveColour == colour {
			points = common.MinusInf
		} else {
			points = common.PlusInf
		}
		r = points
	} else if 0 == depth {
		wp, bp := 0, 0
		for _, cells := range pos.Board {
			for _, cell := range cells {
				inc := 0
				switch cell.Type {
				case common.PiecePawn:
					inc = MaterialIncrementPawn
				case common.PieceKnight:
					inc = MaterialIncrementKnight
				case common.PieceBishop:
					inc = MaterialIncrementBishop
				case common.PieceRook:
					inc = MaterialIncrementRook
				case common.PieceQueen:
					inc = MaterialIncrementQueen
				}
				switch cell.Colour {
				case common.ColourWhite:
					wp += inc
				case common.ColourBlack:
					bp += inc
				}
			}
		}
		diff := wp - bp
		if common.ColourBlack == colour {
			diff = diff * -1
		}
		r = diff
	} else {
		ms := board.FindAllAvailableMoves(pos)
		areOwnMovesEvaluated := pos.ActiveColour == colour
		bs := 0
		if areOwnMovesEvaluated {
			bs = common.MinusInf
		} else {
			bs = common.PlusInf
		}

		for _, m := range ms {
			p2 := board.CommitMove(pos, m)
			s := e.evaluate(p2, colour, depth-1)

			if (areOwnMovesEvaluated && s > bs) || (!areOwnMovesEvaluated && s < bs) {
				bs = s
			}
		}
		r = bs
	}

	return r
}
