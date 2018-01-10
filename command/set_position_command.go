package command

import (
	"github.com/tasmanianfox/dingo/board"
	"github.com/tasmanianfox/dingo/common"
)

const (
	// SetPositionModeMovements Reads movements from Movements property
	SetPositionModeMovements = 0
	// SetPositionModeDirect Assigns position directly from Position
	SetPositionModeDirect = 1
)

type SetPositionCommand struct {
	Movements []board.Movement
	NewGame   bool
	Mode      int
	Position  board.Position
}

func (SetPositionCommand) GetType() int {
	return common.СommandSetPosition
}

func (c *SetPositionCommand) AddMovement(m board.Movement) {
	c.Movements = append(c.Movements, m)
}
