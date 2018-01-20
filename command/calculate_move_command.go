package command

import (
	"github.com/tasmanianfox/dingo/common"
)

type CalculateMoveCommand struct {
}

func (CalculateMoveCommand) GetType() int {
	return common.СommandCalculateMove
}
