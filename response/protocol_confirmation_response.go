package response

import (
	"github.com/tasmanianfox/dingo/common"
)

// ProtocolConfirmationResponse is returned immediately after initialization of the protocol
type ProtocolConfirmationResponse struct {
}

// GetType returns code of response
func (ProtocolConfirmationResponse) GetType() int {
	return common.ResponseProtocolConfirmation
}
