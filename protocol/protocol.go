package protocol

import (
	"github.com/tasmanianfox/dingo/command"
)

type Protocol interface {
	readCommand() command.Command
}
