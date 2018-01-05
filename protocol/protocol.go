package protocol

import (
	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/response"
)

type Protocol interface {
	ReadCommand() (command.Command, bool)
	Output(response.Response)
}
