package main

import (
	"io"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	var r io.Reader = strings.NewReader("uci\ntest\ntest123\nquit\n")
	Run(r)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	r = strings.NewReader("wrong_protocol\ntest\ntest123\nquit\n")
	Run(r)
}
