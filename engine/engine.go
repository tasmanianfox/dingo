package engine

import (
	"github.com/tasmanianfox/dingo/board"
	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/common"
	"github.com/tasmanianfox/dingo/response"
)

type Engine struct {
	Position board.Position
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
	default:
		s = false
	}

	return r, s
}
