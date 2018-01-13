package command

import (
	"github.com/tasmanianfox/dingo/board"
	"github.com/tasmanianfox/dingo/common"
)

type SetPositionCommand struct {
	Position board.Position
}

func (SetPositionCommand) GetType() int {
	return common.СommandSetPosition
}
