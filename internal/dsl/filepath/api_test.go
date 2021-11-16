package filepath_test

import (
	"io/ioutil"
	"testing"

	lua "github.com/yuin/gopher-lua"

	"github.com/vadv/oh-my-pg-linter/internal/dsl/filepath"
)

func TestApi(t *testing.T) {
	data, err := ioutil.ReadFile("./test/test_api.lua")
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}
	state := lua.NewState()
	filepath.Preload(state)
	if err := state.DoString(string(data)); err != nil {
		t.Fatalf("execute test: %s\n", err.Error())
	}
}
