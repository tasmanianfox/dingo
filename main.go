package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/common"
	"github.com/tasmanianfox/dingo/protocol"
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
		p = new(protocol.UciProtocol)
	}
	_, ok := p.(protocol.Protocol)
	if !ok {
		panic("Could not determine the protocol")
	}

	var c command.Command = nil
	for !((c != nil) && (c.GetType() == common.СommandQuit)) {
		s.Scan()
		input := s.Text()

		fmt.Println("In: " + input)
	}
}
