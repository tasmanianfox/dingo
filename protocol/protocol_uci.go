package protocol

import (
	"bufio"
	"os"
	"strings"

	"github.com/tasmanianfox/dingo/command"
)

type UciProtocol struct {
}

func (UciProtocol) readCommand() command.Command {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	var c command.Command
	switch strings.TrimRight(input, "\n") {
	case "isready":
		c = new(command.CheckIfIsEngineReadyCommand)
	case "quit":
		c = new(command.QuitCommand)
	}

	return c
}
