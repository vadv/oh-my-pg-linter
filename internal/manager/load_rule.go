package manager

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

const (
	checkFile   = `check.lua`
	messageFile = `message.md`
	testFile    = `test.lua`
)

func loadRuleToManager(m *manager, b box, path, ruleName string) error {
	// check
	checkData, errReadCheck := b.ReadFile(filepath.Join(path, ruleName, checkFile))
	if errReadCheck != nil {
		return errReadCheck
	}
	if errDo := m.state.DoString(string(checkData)); errDo != nil {
		return fmt.Errorf("load rule %s: %w", ruleName, errDo)
	}
	luaVal := m.state.Get(-1)
	if fun, ok := luaVal.(*lua.LFunction); ok {
		m.rules[ruleName] = fun
	} else {
		return fmt.Errorf("load rule %s: return must be a function", ruleName)
	}
	// messages
	messageFilename := filepath.Join(path, ruleName, messageFile)
	if errStat := b.Stat(messageFilename); errStat == nil {
		data, errReadMessage := b.ReadFile(messageFilename)
		if errReadMessage != nil {
			return errReadMessage
		}
		m.messages[ruleName] = data
	}
	// tests
	testFilename := filepath.Join(path, ruleName, testFile)
	m.tests[ruleName] = m.state.NewTable()
	if errStat := b.Stat(testFilename); errors.Is(errStat, os.ErrNotExist) {
		return nil
	}
	testData, errReadTest := b.ReadFile(testFilename)
	if errReadTest != nil {
		return errReadTest
	}
	if errDo := m.state.DoString(string(testData)); errDo != nil {
		return fmt.Errorf("load test of rule %s: %w", ruleName, errDo)
	}
	luaVal = m.state.Get(-1)
	if table, ok := luaVal.(*lua.LTable); ok {
		m.tests[ruleName] = table
	} else {
		return fmt.Errorf("load test of rule %s: return must be a table", ruleName)
	}
	return nil
}
