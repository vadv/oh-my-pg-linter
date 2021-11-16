package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vadv/oh-my-pg-linter/internal/rules"
)

// Test ...
func Test() *cobra.Command {
	result := &cobra.Command{}
	result.Use = "test [rule]"
	result.Short = "Test rule."
	result.Args = cobra.ExactArgs(1)
	result.Run = func(cmd *cobra.Command, args []string) {
		manager, errManager := rules.New(viper.GetString("rules"))
		if errManager != nil {
			log.Fatal(fmt.Errorf("load manger: %w", errManager))
		}
		if err := manager.Test(args[0]); err != nil {
			log.Fatal(err)
		}
	}
	return result
}
