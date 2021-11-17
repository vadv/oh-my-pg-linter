// Package commands ...
package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vadv/oh-my-pg-linter/internal/rules"
)

// Check ...
func Check() *cobra.Command {
	result := &cobra.Command{}
	result.Use = "check [glob ...]"
	result.Short = "Check sql-files."
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
					var q string
					if check.Query() != nil {
						q = strings.Trim(strings.Trim(*check.Query(), " "), "\n")
					}
					data := markdown.Render(fmt.Sprintf("%s\n\trule: `%s`\n\tquery: `%s`\n%s\n",
						f, r, q, check.Message()), 120, 2)
					fmt.Printf("%s\n", data)
					errorCount++
				}
			}
		}
		if errorCount != 0 {
			fmt.Printf("%s\n", strings.Repeat("-", 20))
			fmt.Printf("Found %d errors.\n", errorCount)
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
