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
	var varExcludeRules string
	result.Flags().StringVar(&varExcludeRules, "exclude", "", "Exclude rules (delimited by `,`).")
	result.Run = func(cmd *cobra.Command, args []string) {
		manager := rules.New()
		if errManager := addRuleDirs(manager, viper.GetString("rules")); errManager != nil {
			log.Fatal(fmt.Errorf("load manger: %w", errManager))
		}
		files, errGlob := getListOfFiles(args)
		if errGlob != nil {
			log.Fatal(fmt.Errorf("glob files: %w", errGlob))
		}
		excludeRules := strings.Split(varExcludeRules, ",")
		var errorCount int
		for _, f := range files {
			for _, r := range manager.ListRules() {
				if skipRule(r, excludeRules) {
					continue
				}
				check, errCheck := manager.Check(f, r)
				if errCheck != nil {
					log.Fatal(errCheck)
				}
				if !check.Passed() {
					var q string
					if check.Query() != nil {
						q = strings.Trim(strings.Trim(*check.Query(), " "), "\n")
					}
					data := markdown.Render(fmt.Sprintf(""+
						"# File\n[%s](%s)\n"+
						"# Rule\n`%s`\n"+
						"# Statement\n"+
						"```sql\n"+
						"%s\n"+
						"```"+
						"\n%s\n",
						filepath.Base(f), f,
						r,
						q,
						check.Message()), 140, 2)
					fmt.Printf("%s\n", data)
					errorCount++
				}
			}
		}
		if errorCount != 0 {
			fmt.Printf("%s\n", strings.Repeat("-", 20))
			fmt.Printf("Found %d error(s).\n", errorCount)
			os.Exit(errorCount)
		}
	}
	return result
}

func getListOfFiles(args []string) ([]string, error) {
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

func skipRule(rule string, exclude []string) bool {
	for _, e := range exclude {
		if e == rule {
			return true
		}
	}
	return false
}
