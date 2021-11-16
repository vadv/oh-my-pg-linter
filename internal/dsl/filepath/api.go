// Package filepath implements golang filepath functionality for lua.
package filepath

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

// Basename lua filepath.basename(path) returns the last element of path.
func Basename(L *lua.LState) int {
	path := L.CheckString(1)
	L.Push(lua.LString(filepath.Base(path)))
	return 1
}

// Dir lua filepath.dir(path) returns all but the last element of path, typically the path's directory.
func Dir(L *lua.LState) int {
	path := L.CheckString(1)
	L.Push(lua.LString(filepath.Dir(path)))
	return 1
}

// Ext lua filepath.ext(path) returns the file name extension used by path.
func Ext(L *lua.LState) int {
	path := L.CheckString(1)
	L.Push(lua.LString(filepath.Ext(path)))
	return 1
}

// Join lua fileapth.join(path, ...) joins any number of path elements into a single path, adding a Separator if necessary.
func Join(L *lua.LState) int {
	path := L.CheckString(1)
	for i := 2; i <= L.GetTop(); i++ {
		add := L.CheckAny(i).String()
		path = filepath.Join(path, add)
	}
	L.Push(lua.LString(path))
	return 1
}

// Separator lua filepath.separator() OS-specific path separator.
func Separator(L *lua.LState) int {
	L.Push(lua.LString(filepath.Separator))
	return 1
}

// ListSeparator lua filepath.list_separator() OS-specific path list separator.
func ListSeparator(L *lua.LState) int {
	L.Push(lua.LString(filepath.ListSeparator))
	return 1
}

// Glob lua filepath.glob(pattern) returns the names of all files matching pattern or nil if there is no matching file.
func Glob(L *lua.LState) int {
	pattern := L.CheckString(1)
	files, err := filepath.Glob(pattern)
	if err != nil {
		L.Push(lua.LNil)
		return 1
	}
	result := L.CreateTable(len(files), 0)
	for _, file := range files {
		result.Append(lua.LString(file))
	}
	L.Push(result)
	return 1
}

// WorkDir lua path.work_dir() returns path and err.
func WorkDir(L *lua.LState) int {
	currDir, err := os.Getwd()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(currDir))
	return 1
}

// ScriptDir lua path.script_dir() returns path and err.
func ScriptDir(L *lua.LState) int {
	path := L.Where(1)
	if path == "" {
		L.Push(lua.LNil)
		L.Push(lua.LString("source file don't detected correctly"))
		return 2
	}
	// ./plugins/pgss2ch/test.lua:2:  => ./plugins/pgss2ch/test.lua
	path = strings.Split(path, ":")[0]
	path = filepath.Clean(path)
	path = filepath.Dir(path)
	L.Push(lua.LString(path))
	return 1
}

// Exists lua path.exists() returns bool and err.
func Exists(L *lua.LState) int {
	path := L.CheckString(1)
	_, err := os.Stat(path)
	if err == nil {
		L.Push(lua.LBool(true))
		return 1
	}
	if errors.Is(err, os.ErrNotExist) {
		L.Push(lua.LBool(false))
		return 1
	}
	L.Push(lua.LNil)
	L.Push(lua.LString(err.Error()))
	return 2
}
