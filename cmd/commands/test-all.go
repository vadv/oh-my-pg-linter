package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vadv/oh-my-pg-linter/internal/manager"
)

// TestAll ...
func TestAll() *cobra.Command {
	result := &cobra.Command{}
	result.Use = "test-all"
	result.Short = "Tests all rules."
	result.Run = func(cmd *cobra.Command, args []string) {
		m := manager.New()
		if ruleDirs := viper.GetString("rules"); ruleDirs != "" {
			if errManager := addRuleDirs(m, ruleDirs); errManager != nil {
				log.Fatal(fmt.Errorf("load rule: %w", errManager))
			}
		}
		for _, r := range m.ListRules() {
			if err := m.Test(r); err != nil {
				log.Fatalf("test rule %s: %s\n", r, err.Error())
			}
		}
	}
	return result
}
