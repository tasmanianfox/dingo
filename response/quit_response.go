package response

import (
	"github.com/tasmanianfox/dingo/common"
)

type QuitResponse struct {
}

func (QuitResponse) GetType() int {
	return common.ResponseQuit
}
