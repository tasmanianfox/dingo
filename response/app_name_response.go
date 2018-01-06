package response

import (
	"github.com/tasmanianfox/dingo/common"
)

type AppNameResponse struct {
}

func (AppNameResponse) GetType() int {
	return common.ResponseAppName
}
