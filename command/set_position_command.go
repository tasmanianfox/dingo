package command

import (
	"github.com/tasmanianfox/dingo/board"
	"github.com/tasmanianfox/dingo/common"
)

const (
	SetPositionModeMoves = 0
	SetPositionModeFen   = 1
)

type SetPositionCommand struct {
	Movements []board.Movement
	NewGame   bool
	Mode      int
}

func (SetPositionCommand) GetType() int {
	return common.СommandSetPosition
}

func (c *SetPositionCommand) AddMovement(m board.Movement) {
	c.Movements = append(c.Movements, m)
}
