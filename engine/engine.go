package engine

import (
	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/common"
	"github.com/tasmanianfox/dingo/response"
)

type Engine struct {
}

func (Engine) HandleCommand(c command.Command) response.Response {
	var r response.Response
	switch c.GetType() {
	case common.СommandQuit:
		r = new(response.QuitResponse)
	}

	return r
}
