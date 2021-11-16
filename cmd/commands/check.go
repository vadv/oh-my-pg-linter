// Package commands ...
package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vadv/oh-my-pg-linter/internal/rules"
)

// Check ...
func Check() *cobra.Command {
	result := &cobra.Command{}
	result.Use = "check [glob ...]"
	result.Short = "Check files with migrations."
	result.Run = func(cmd *cobra.Command, args []string) {
		manager, errManager := rules.New(viper.GetString("rules"))
		if errManager != nil {
			log.Fatal(fmt.Errorf("load manger: %w", errManager))
		}
		files, errGlob := getListFiles(args)
		if errGlob != nil {
			log.Fatal(fmt.Errorf("glob files: %w", errGlob))
		}
		var errorCount int
		for _, f := range files {
			for _, r := range manager.ListRules() {
				check, errCheck := manager.Check(f, r)
				if errCheck != nil {
					log.Fatal(errCheck)
				}
				if !check.Passed() {
					message := string(check.Message())
					fmt.Printf("%s: check rule %s:\n%s",
						f, r, message)
					errorCount++
				}
			}
		}
		if errorCount != 0 {
			os.Exit(errorCount)
		}
	}
	return result
}

func getListFiles(args []string) ([]string, error) {
	result := make([]string, 0)
	for _, arg := range args {
		files, err := filepath.Glob(arg)
		if err != nil {
			return nil, err
		}
		result = append(result, files...)
	}
	return result, nil
}
