package commands

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vadv/oh-my-pg-linter/internal/dsl/filepath"
	"github.com/vadv/oh-my-pg-linter/internal/dsl/inspect"
	"github.com/vadv/oh-my-pg-linter/internal/dsl/json"
	"github.com/vadv/oh-my-pg-linter/internal/dsl/parser"
	lua "github.com/yuin/gopher-lua"
)

// Run ...
func Run() *cobra.Command {
	result := &cobra.Command{}
	result.Use = "run [file]"
	result.Short = "Run lua file."
	result.Args = cobra.ExactArgs(1)
	result.Run = func(cmd *cobra.Command, args []string) {
		state := lua.NewState()
		filepath.Preload(state)
		inspect.Preload(state)
		json.Preload(state)
		parser.Preload(state)
		if err := state.DoFile(args[0]); err != nil {
			log.Fatal(err)
		}
	}
	return result
}
