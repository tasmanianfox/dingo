package response

import (
	"github.com/tasmanianfox/dingo/common"
)

// EmptyResponse represents a response that does not require any output
type EmptyResponse struct {
}

// GetType returns code of response
func (EmptyResponse) GetType() int {
	return common.ResponseEmpty
}
