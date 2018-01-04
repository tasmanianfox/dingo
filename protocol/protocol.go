package protocol

import "github.com/tasmanianfox/dingo/command"

type Protocol interface {
	ReadCommand() command.Command
}
