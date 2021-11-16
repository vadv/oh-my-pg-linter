package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vadv/oh-my-pg-linter/internal/rules"
)

// TestAll ...
func TestAll() *cobra.Command {
	result := &cobra.Command{}
	result.Use = "test-all"
	result.Short = "Tests all rules."
	result.Run = func(cmd *cobra.Command, args []string) {
		manager, errManager := rules.New(viper.GetString("rules"))
		if errManager != nil {
			log.Fatal(fmt.Errorf("load manger: %w", errManager))
		}
		for _, r := range manager.ListRules() {
			if err := manager.Test(r); err != nil {
				log.Fatalf("test rule %s: %s\n", r, err.Error())
			}
		}
	}
	return result
}
