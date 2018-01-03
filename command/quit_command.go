package command

import (
	"github.com/tasmanianfox/dingo/common"
)

type QuitCommand struct {
}

func (QuitCommand) GetType() int {
	return common.СommandQuit
}
