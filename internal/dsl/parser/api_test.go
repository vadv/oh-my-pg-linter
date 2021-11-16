package parser_test

import (
	"testing"

	"github.com/vadv/oh-my-pg-linter/internal/dsl/inspect"
	"github.com/vadv/oh-my-pg-linter/internal/dsl/parser"
	lua "github.com/yuin/gopher-lua"
)

func TestApi(t *testing.T) {
	state := lua.NewState()
	inspect.Preload(state)
	parser.Preload(state)
	if err := state.DoFile("./test/test_api.lua"); err != nil {
		t.Fatalf("execute test: %s\n", err.Error())
	}
}
