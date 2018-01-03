package main

import (
	"io"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	var r io.Reader = strings.NewReader("uci\ntest\ntest123\nquit\n")
	Run(r)
}
