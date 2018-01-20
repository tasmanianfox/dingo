package main

import (
	"bufio"
	"io"
	"os"

	"github.com/tasmanianfox/dingo/common"
	"github.com/tasmanianfox/dingo/engine"
	"github.com/tasmanianfox/dingo/protocol"
	"github.com/tasmanianfox/dingo/response"
)

func main() {
	var r io.Reader = os.Stdin
	var w io.Writer = os.Stdout
	Run(r, w)
}

func Run(r io.Reader, w io.Writer) {
	var p protocol.Protocol = nil
	s := bufio.NewScanner(r)
	bw := bufio.NewWriter(w)
	s.Scan()
	if "uci" == s.Text() {
		p = protocol.NewUciProtocol(s, bw)
	}
	_, ok := p.(protocol.Protocol)
	if !ok {
		panic("Could not determine the protocol")
	}

	ch := make(chan response.Response)
	e := engine.Engine{ResponseChannel: ch}
	go ReadResponseAsync(p, ch)
	for {
		c, success := p.ReadCommand()
		if false == success {
			continue
		}

		response, success := e.HandleCommand(c)
		if false == success {
			continue
		} else if (response != nil) && (common.ResponseQuit == response.GetType()) {
			break
		}
		p.Output(response)
	}
}

func ReadResponseAsync(p protocol.Protocol, ch chan response.Response) {
	for {
		r := <-ch
		p.Output(r)
	}
}
