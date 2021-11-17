package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vadv/oh-my-pg-linter/internal/manager"
)

// Test ...
func Test() *cobra.Command {
	result := &cobra.Command{}
	result.Use = "test [rule]"
	result.Short = "Test rule."
	result.Args = cobra.ExactArgs(1)
	result.Run = func(cmd *cobra.Command, args []string) {
		m := manager.New()
		if ruleDirs := viper.GetString("rules"); ruleDirs != "" {
			if errManager := addRuleDirs(m, ruleDirs); errManager != nil {
				log.Fatal(fmt.Errorf("load rule: %w", errManager))
			}
		}
		if err := m.Test(args[0]); err != nil {
			log.Fatal(err)
		}
	}
	return result
}
