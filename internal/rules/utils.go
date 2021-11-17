package rules

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/vadv/oh-my-pg-linter/internal/dsl/parser"
	lua "github.com/yuin/gopher-lua"
)

func loadRuleToManager(m *manager, dir, ruleName string) error {
	// check
	checkData, errReadCheck := ioutil.ReadFile(filepath.Join(dir, ruleName, "check.lua"))
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
	messageFilename := filepath.Join(dir, ruleName, "message.md")
	if _, errStat := os.Stat(messageFilename); errStat == nil {
		data, errReadMessage := ioutil.ReadFile(messageFilename)
		if errReadMessage != nil {
			return errReadMessage
		}
		m.messages[ruleName] = data
	}
	// tests
	testFilename := filepath.Join(dir, ruleName, "test.lua")
	m.tests[ruleName] = m.state.NewTable()
	if _, errStat := os.Stat(testFilename); errors.Is(errStat, os.ErrNotExist) {
		return nil
	}
	testData, errReadTest := ioutil.ReadFile(testFilename)
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

func (m *manager) runRule(content, rule string) (*response, error) {
	ruleFunc, okFunc := m.rules[rule]
	if !okFunc {
		return nil, fmt.Errorf("rule not found")
	}
	if errCall := m.state.CallByParam(lua.P{
		Fn:      m.loadQuery,
		NRet:    -1,
		Protect: true,
	}, lua.LString(content)); errCall != nil {
		return nil, fmt.Errorf("parse tree: %w", errCall)
	}
	luaVal := m.state.Get(-1)
	var tree *lua.LTable
	if data, ok := luaVal.(*lua.LTable); ok {
		tree = data
	} else {
		return nil, fmt.Errorf("can't parse tree: %v", luaVal)
	}
	if err := m.state.CallByParam(lua.P{
		Fn:      ruleFunc,
		NRet:    -1,
		Protect: true,
	}, tree); err != nil {
		return nil, err
	}
	return m.rebuke(m.state.Get(-1), rule)
}

func (m *manager) rebuke(luaVal lua.LValue, rule string) (*response, error) {
	result := &response{passed: true}
	tbl, okTbl := luaVal.(*lua.LTable)
	if !okTbl {
		return nil, fmt.Errorf("unexcepted type: %#v", luaVal)
	}
	var err error
	tbl.ForEach(func(_, v lua.LValue) {
		ud, okUd := v.(*lua.LUserData)
		if !okUd {
			err = fmt.Errorf("unexcepted type element of table: %#v, required userdata", v)
			return
		}
		val := ud.Value
		stmt, okStmt := val.(*parser.Stmt)
		if !okStmt {
			err = fmt.Errorf("unexcepted type element of table: %#v, required stmt", val)
			return
		}
		if !stmt.IsNoLint(rule) {
			result.passed = false
			result.query = stmt.Query()
			result.message = m.messages[rule]
			return
		}
	})
	return result, err
}
