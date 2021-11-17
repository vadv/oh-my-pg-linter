// Package parser ...
package parser

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v2"
	"github.com/vadv/oh-my-pg-linter/internal/dsl/json"
	lua "github.com/yuin/gopher-lua"
)

const (
	userDataNameStmt = `user_data_stmt`
)

// Stmt ...
type Stmt struct {
	query   string
	json    lua.LValue
	noLints []string
}

// NoLints ...
func (s *Stmt) NoLints() []string {
	return s.noLints
}

// IsNoLint ...
func (s *Stmt) IsNoLint(rule string) bool {
	for _, x := range s.noLints {
		if rule == x {
			return true
		}
	}
	return false
}

// Query ...
func (s *Stmt) Query() string {
	return s.query
}

func checkStmt(L *lua.LState, n int) *Stmt {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.(*Stmt); ok {
		return v
	}
	L.ArgError(n, "ud expected")
	return nil
}

// Query ...
func Query(L *lua.LState) int {
	stmt := checkStmt(L, 1)
	L.Push(lua.LString(stmt.query))
	return 1
}

// Tree ...
func Tree(L *lua.LState) int {
	stmt := checkStmt(L, 1)
	L.Push(stmt.json)
	return 1
}

// IsNoLint ud:is_no_lint(rulename) returns bool.
func IsNoLint(L *lua.LState) int {
	stmt := checkStmt(L, 1)
	val := L.CheckString(2)
	L.Push(lua.LBool(stmt.IsNoLint(val)))
	return 1
}

// Parse lua parser.tree(string) returns (user_data, err).
func Parse(L *lua.LState) int {
	str := L.CheckString(1)
	tree, errParse := pg_query.Parse(str)
	if errParse != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(errParse.Error()))
		return 2
	}
	result := L.NewTable()
	for _, s := range tree.Stmts {
		query := str[int(s.StmtLocation):int(s.StmtLen+s.StmtLocation)]
		tree, errParse := pg_query.ParseToJSON(query)
		if errParse != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(errParse.Error()))
			return 2
		}
		jsonVal, errDecode := json.ValueDecode(L, []byte(tree))
		if errDecode != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(errDecode.Error()))
			return 2
		}
		// jsonVal = {stmts = { }}
		tbl, ok := jsonVal.(*lua.LTable)
		if !ok {
			L.Push(lua.LNil)
			L.Push(lua.LString(fmt.Sprintf("parsing of %#v must be table", query)))
			return 2
		}
		jsonVal = tbl.RawGetString("stmts")
		// { {} }
		tbl, ok = jsonVal.(*lua.LTable)
		if !ok {
			L.Push(lua.LNil)
			L.Push(lua.LString(fmt.Sprintf("parsing of %#v must be table with stmts", query)))
			return 2
		}
		jsonVal = tbl.RawGetInt(1)
		// {stmt = {}}
		tbl, ok = jsonVal.(*lua.LTable)
		if !ok {
			L.Push(lua.LNil)
			L.Push(lua.LString(fmt.Sprintf("parsing of %#v must be table with stmt", query)))
			return 2
		}
		l := &Stmt{json: tbl.RawGetString("stmt"), query: query, noLints: NoLintParse(query)}
		ud := L.NewUserData()
		ud.Value = l
		L.SetMetatable(ud, L.GetTypeMetatable(userDataNameStmt))
		result.Append(ud)
	}
	L.Push(result)
	return 1
}
