package response

import (
	"github.com/tasmanianfox/dingo/common"
)

type ReadyResponse struct {
}

func (ReadyResponse) GetType() int {
	return common.ResponseReady
}
