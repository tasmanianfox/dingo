package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/common"
	"github.com/tasmanianfox/dingo/protocol"
)

func main() {
	var p protocol.Protocol = nil
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if "uci" == strings.TrimRight(input, "\n") {
		p = new(protocol.UciProtocol)
	}
	_, ok := p.(protocol.Protocol)
	if !ok {
		panic("Could not determine the protocol")
	}

	var c command.Command = nil
	for (c != nil) && (c.GetType() == common.СommandQuit) {
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")
	}
}
