package manager

import (
	"fmt"

	"github.com/vadv/oh-my-pg-linter/internal/dsl/parser"
	lua "github.com/yuin/gopher-lua"
)

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
		return nil, fmt.Errorf("rule %s: parse tree: %w", rule, errCall)
	}
	luaVal := m.state.Get(-1)
	var tree *lua.LTable
	if data, ok := luaVal.(*lua.LTable); ok {
		tree = data
	} else {
		return nil, fmt.Errorf("rule %s: can't parse tree: %v", rule, luaVal)
	}
	if err := m.state.CallByParam(lua.P{
		Fn:      ruleFunc,
		NRet:    -1,
		Protect: true,
	}, tree); err != nil {
		return nil, fmt.Errorf("run rule %s: %w", rule, err)
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
