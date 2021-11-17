package commands

import (
	"path/filepath"

	"github.com/vadv/oh-my-pg-linter/internal/rules"
)

func addRuleDirs(m rules.Manager, dirs string) error {
	for _, d := range filepath.SplitList(dirs) {
		if err := m.AddRuleDir(d); err != nil {
			return err
		}
	}
	return nil
}
