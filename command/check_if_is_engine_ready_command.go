package command

import (
	"github.com/tasmanianfox/dingo/common"
)

type CheckIfIsEngineReadyCommand struct {
}

func (CheckIfIsEngineReadyCommand) GetType() int {
	return common.СommandCheckIfIsEngineReady
}
