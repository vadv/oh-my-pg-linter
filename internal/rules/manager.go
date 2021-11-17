// Package rules ...
package rules

import (
	// embed files.
	_ "embed"
	"fmt"
	"io/ioutil"
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

var (
	//go:embed parse_tree.lua
	getTreeFromFileInLua string
)

type manager struct {
	directory string
	state     *lua.LState
	rules     map[string]*lua.LFunction
	tests     map[string]*lua.LTable
	messages  map[string][]byte
	loadQuery *lua.LFunction
}

// New returns new Manager.
func New() Manager {
	m := &manager{state: NewState()}
	if errLoad := m.state.DoString(getTreeFromFileInLua); errLoad != nil {
		panic(errLoad)
	}
	luaVal := m.state.Get(-1)
	if fun, ok := luaVal.(*lua.LFunction); ok {
		m.loadQuery = fun
	} else {
		panic(fmt.Errorf("can't parse getTreeFromFileInLua: %v", luaVal))
	}
	m.rules = make(map[string]*lua.LFunction)
	m.tests = make(map[string]*lua.LTable)
	m.messages = make(map[string][]byte)
	return m
}

func (m *manager) AddRuleDir(dir string) error {
	dir = filepath.Clean(dir)
	files, errRead := ioutil.ReadDir(dir)
	if errRead != nil {
		return errRead
	}
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		if err := loadRuleToManager(m, dir, f.Name()); err != nil {
			return err
		}
	}
	return nil
}

func (m *manager) ListRules() []string {
	result := make([]string, 0)
	for r := range m.rules {
		result = append(result, r)
	}
	return result
}

func (m *manager) Check(file, rule string) (Response, error) {
	fileContent, errReadFileContent := ioutil.ReadFile(file)
	if errReadFileContent != nil {
		return nil, errReadFileContent
	}
	return m.runRule(string(fileContent), rule)
}

func (m *manager) Test(rule string) error {
	t := m.tests[rule]
	if t == nil {
		return fmt.Errorf("test for rule `%s` is not found", rule)
	}
	var err error
	t.ForEach(func(_, v lua.LValue) {
		tbl, ok := v.(*lua.LTable)
		if !ok {
			err = fmt.Errorf("value is not table")
			return
		}
		sql := string(tbl.RawGetString("sql").(lua.LString))
		mustPassed := bool(tbl.RawGetString("passed").(lua.LBool))
		if len(sql) > 0 && sql[:len(sql)-1] != ";" {
			sql = sql + ";"
		}
		resp, errRunRule := m.runRule(sql, rule)
		if errRunRule != nil {
			err = errRunRule
			return
		}
		if resp.Passed() != mustPassed {
			err = fmt.Errorf("sql: %#v get: %v except: %v", sql, resp.Passed(), mustPassed)
			return
		}
	})
	return err
}
