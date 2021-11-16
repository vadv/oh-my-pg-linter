package json_test

import (
	"testing"

	lua "github.com/yuin/gopher-lua"

	"github.com/vadv/oh-my-pg-linter/internal/dsl/inspect"
	"github.com/vadv/oh-my-pg-linter/internal/dsl/json"
)

func TestApi(t *testing.T) {
	state := lua.NewState()
	json.Preload(state)
	inspect.Preload(state)
	if err := state.DoFile("./test/test_api.lua"); err != nil {
		t.Fatalf("execute test: %s\n", err.Error())
	}
}
