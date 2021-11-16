// Package rules ...
package rules

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	filepath2 "path/filepath"

	"github.com/vadv/oh-my-pg-linter/internal/dsl/filepath"
	"github.com/vadv/oh-my-pg-linter/internal/dsl/inspect"
	"github.com/vadv/oh-my-pg-linter/internal/dsl/json"
	"github.com/vadv/oh-my-pg-linter/internal/dsl/parser"
	lua "github.com/yuin/gopher-lua"
)

const (
	preloadLibInLua = `
filepath = require("filepath")
inspect = require("inspect")
json = require("json")
parser = require("parser")
`
	getTreeFromFileInLua = `
local function parseTree(content)
	local parser = require("parser")
	local result, err = parser.parse(content)
	if err then error(err) end
	return result
end

return parseTree 
`
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
func New(dir string) (Manager, error) {
	dir = filepath2.Clean(dir)
	m := &manager{directory: dir}
	state := lua.NewState()
	filepath.Preload(state)
	inspect.Preload(state)
	json.Preload(state)
	parser.Preload(state)
	if errLoad := state.DoString(preloadLibInLua); errLoad != nil {
		panic(errLoad)
	}
	if errLoad := state.DoString(getTreeFromFileInLua); errLoad != nil {
		panic(errLoad)
	}
	luaVal := state.Get(-1)
	if fun, ok := luaVal.(*lua.LFunction); ok {
		m.loadQuery = fun
	} else {
		panic(fmt.Errorf("can't parse getTreeFromFileInLua: %v", luaVal))
	}
	files, errRead := ioutil.ReadDir(dir)
	if errRead != nil {
		return nil, errRead
	}
	rules := make(map[string]*lua.LFunction)
	tests := make(map[string]*lua.LTable)
	messages := make(map[string][]byte)
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		ruleName := f.Name()
		// check
		checkData, errReadCheck := ioutil.ReadFile(filepath2.Join(dir, ruleName, "check.lua"))
		if errReadCheck != nil {
			return nil, errReadCheck
		}
		if errDo := state.DoString(string(checkData)); errDo != nil {
			return nil, fmt.Errorf("load rule %s: %w", ruleName, errDo)
		}
		luaVal := state.Get(-1)
		if fun, ok := luaVal.(*lua.LFunction); ok {
			rules[ruleName] = fun
		} else {
			return nil, fmt.Errorf("load rule %s: return must be a function", ruleName)
		}
		// tests
		testFilename := filepath2.Join(dir, ruleName, "test.lua")
		tests[ruleName] = state.NewTable()
		if _, errStat := os.Stat(testFilename); errors.Is(errStat, os.ErrNotExist) {
			continue
		}
		testData, errReadTest := ioutil.ReadFile(testFilename)
		if errReadTest != nil {
			return nil, errReadTest
		}
		if errDo := state.DoString(string(testData)); errDo != nil {
			return nil, fmt.Errorf("load test of rule %s: %w", ruleName, errDo)
		}
		luaVal = state.Get(-1)
		if table, ok := luaVal.(*lua.LTable); ok {
			tests[ruleName] = table
		} else {
			return nil, fmt.Errorf("load test of rule %s: return must be a table", ruleName)
		}
	}
	m.state = state
	m.rules = rules
	m.tests = tests
	m.messages = messages
	return m, nil
}

func (m *manager) ListRules() []string {
	result := make([]string, 0)
	for r := range m.rules {
		result = append(result, r)
	}
	return result
}

func (m *manager) Check(file, rule string) (Response, error) {
	ruleFunc, okFunc := m.rules[rule]
	if !okFunc {
		return nil, fmt.Errorf("rule not found")
	}
	fileContent, errReadFileContent := ioutil.ReadFile(file)
	if errReadFileContent != nil {
		return nil, errReadFileContent
	}
	if errCall := m.state.CallByParam(lua.P{
		Fn:      m.loadQuery,
		NRet:    -1,
		Protect: true,
	}, lua.LString(fileContent)); errCall != nil {
		return nil, fmt.Errorf("load file %s: parse tree: %w", file, errCall)
	}
	luaVal := m.state.Get(-1)
	var tree *lua.LTable
	if data, ok := luaVal.(*lua.LTable); ok {
		tree = data
	} else {
		panic(fmt.Errorf("can't parse tree: %v", luaVal))
	}
	if err := m.state.CallByParam(lua.P{
		Fn:      ruleFunc,
		NRet:    -1,
		Protect: true,
	}, tree); err != nil {
		return nil, err
	}
	luaVal = m.state.Get(-1)
	result := &response{}
	if res, ok := luaVal.(lua.LString); ok {
		result.message = []byte(res)
	} else if luaVal.Type() == lua.LTNil {
		result.passed = true
	} else {
		return nil, fmt.Errorf("unexcepted type of result check function: %#v", luaVal)
	}
	return result, nil
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
		if errCall := m.state.CallByParam(lua.P{
			Fn:      m.loadQuery,
			NRet:    -1,
			Protect: true,
		}, lua.LString(sql)); errCall != nil {
			panic(fmt.Errorf("load sql %#v: parse tree: %w", sql, errCall))
		}
		luaVal := m.state.Get(-1)
		var tree *lua.LTable
		if data, ok := luaVal.(*lua.LTable); ok {
			tree = data
		} else {
			panic(fmt.Errorf("can't parse tree: %v", luaVal))
		}
		if errRule := m.state.CallByParam(lua.P{
			Fn:      m.rules[rule],
			NRet:    -1,
			Protect: true,
		}, tree); errRule != nil {
			panic(fmt.Errorf("run rule: %w", errRule))
		}
		var passed bool
		luaVal = m.state.Get(-1)
		if _, ok := luaVal.(lua.LString); ok {
			passed = false
		} else if luaVal.Type() == lua.LTNil {
			passed = true
		} else {
			passed = false
			panic(fmt.Errorf("unexcepted type of result check function: %#v", luaVal))
		}
		if passed != mustPassed {
			err = fmt.Errorf("sql: %#v get: %v except: %v", sql, passed, mustPassed)
			return
		}
	})
	return err
}
