package parser

import (
	lua "github.com/yuin/gopher-lua"
)

// Preload adds parser to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//  local parser = require("parser")
func Preload(L *lua.LState) {
	L.PreloadModule("parser", Loader)
}

// Loader is the module loader function.
func Loader(L *lua.LState) int {
	ud := L.NewTypeMetatable(userDataNameStmt)
	L.SetGlobal(userDataNameStmt, ud)
	L.SetField(ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"query":      Query,
		"tree":       Tree,
		"is_no_lint": IsNoLint,
	}))

	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}

var api = map[string]lua.LGFunction{
	"parse": Parse,
}
