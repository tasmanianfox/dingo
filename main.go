package main

import (
	"bufio"
	"io"
	"os"

	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/common"
	"github.com/tasmanianfox/dingo/engine"
	"github.com/tasmanianfox/dingo/protocol"
	"github.com/tasmanianfox/dingo/response"
)

func main() {
	var r io.Reader = os.Stdin
	Run(r)
}

func Run(r io.Reader) {
	var p protocol.Protocol = nil
	s := bufio.NewScanner(r)
	s.Scan()
	if "uci" == s.Text() {
		p = protocol.NewUciProtocol(s)
	}
	_, ok := p.(protocol.Protocol)
	if !ok {
		panic("Could not determine the protocol")
	}

	var e engine.Engine
	var c command.Command = nil
	var response response.Response = nil
	for !((response != nil) && (common.ResponseQuit == response.GetType())) {
		c = p.ReadCommand()
		if nil == c {
			continue
		}
		response = e.HandleCommand(c)
	}
}
