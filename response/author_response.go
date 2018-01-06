package response

import (
	"github.com/tasmanianfox/dingo/common"
)

type AuthorResponse struct {
}

func (AuthorResponse) GetType() int {
	return common.ResponseAuthor
}
