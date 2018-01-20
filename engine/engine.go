package engine

import (
	"math/rand"
	"time"

	"github.com/tasmanianfox/dingo/board"
	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/common"
	"github.com/tasmanianfox/dingo/response"
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
	rand.Seed(time.Now().UnixNano())
	ms := board.FindAllAvailableMoves(e.Position)

	numMs := len(ms)
	m := board.Move{}
	if numMs > 1 {
		m = ms[rand.Intn(len(ms)-1)]
	} else if numMs == 1 {
		m = ms[0]
	} else {
		panic("Cannot find any moves")
	}
	r := response.MoveResponse{Move: m}
	e.ResponseChannel <- r
}
