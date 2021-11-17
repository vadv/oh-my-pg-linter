package commands

import (
	"path/filepath"

	"github.com/vadv/oh-my-pg-linter/internal/manager"
)

func addRuleDirs(m manager.Manager, dirs string) error {
	for _, d := range filepath.SplitList(dirs) {
		if err := m.AddRuleDir(d); err != nil {
			return err
		}
	}
	return nil
}
