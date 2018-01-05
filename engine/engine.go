package engine

import (
	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/common"
	"github.com/tasmanianfox/dingo/response"
)

type Engine struct {
}

func (Engine) HandleCommand(c command.Command) (response.Response, bool) {
	var r response.Response
	var s bool = true
	switch c.GetType() {
	case common.СommandQuit:
		r = new(response.QuitResponse)
	case common.СommandCheckIfIsEngineReady:
		r = new(response.ReadyResponse)
	default:
		s = false
	}

	return r, s
}
