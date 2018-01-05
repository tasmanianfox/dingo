package engine

import (
	"testing"

	"github.com/tasmanianfox/dingo/command"
	"github.com/tasmanianfox/dingo/common"
)

func TestHandleCommand(t *testing.T) {
	var e = new(Engine)
	var c = new(command.QuitCommand)
	var r, s = e.HandleCommand(c)
	if !(true == s && common.ResponseQuit == r.GetType()) {
		t.Errorf("Expected: ResponseQuit")
	}
}
