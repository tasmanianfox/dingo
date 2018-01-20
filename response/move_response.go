package response

import (
	"github.com/tasmanianfox/dingo/board"
	"github.com/tasmanianfox/dingo/common"
)

type MoveResponse struct {
	Move board.Move
}

func (MoveResponse) GetType() int {
	return common.ResponseMove
}
