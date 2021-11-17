package commands

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vadv/oh-my-pg-linter/internal/manager"
)

// Run ...
func Run() *cobra.Command {
	result := &cobra.Command{}
	result.Use = "run [file]"
	result.Short = "Run lua file."
	result.Args = cobra.ExactArgs(1)
	result.Run = func(cmd *cobra.Command, args []string) {
		state := manager.NewState()
		if err := state.DoFile(args[0]); err != nil {
			log.Fatal(err)
		}
	}
	return result
}
