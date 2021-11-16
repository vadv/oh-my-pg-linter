//go:build !windows && !plan9
// +build !windows,!plan9

package filepath_test

import (
	"log"

	lua "github.com/yuin/gopher-lua"

	"github.com/vadv/oh-my-pg-linter/internal/dsl/filepath"
	"github.com/vadv/oh-my-pg-linter/internal/dsl/inspect"
)

// filepath.ext(string).
func ExampleExt() {
	state := lua.NewState()
	filepath.Preload(state)
	source := `
    local filepath = require("filepath")
    local result = filepath.ext("/var/tmp/file.name")
    print(result)
`
	if err := state.DoString(source); err != nil {
		log.Fatal(err.Error())
	}
	// Output:
	// .name
}

// filepath.basename(string).
func ExampleBasename() {
	state := lua.NewState()
	filepath.Preload(state)
	source := `
    local filepath = require("filepath")
    local result = filepath.basename("/var/tmp/file.name")
    print(result)
`
	if err := state.DoString(source); err != nil {
		log.Fatal(err.Error())
	}
	// Output:
	// file.name
}

// filepath.basename(string).
func ExampleDir() {
	state := lua.NewState()
	filepath.Preload(state)
	source := `
    local filepath = require("filepath")
    local result = filepath.dir("/var/tmp/file.name")
    print(result)
`
	if err := state.DoString(source); err != nil {
		log.Fatal(err.Error())
	}
	// Output:
	// /var/tmp
}

// filepath.basename(string).
func ExampleJoin() {
	state := lua.NewState()
	filepath.Preload(state)
	source := `
    local filepath = require("filepath")
    local result = filepath.join("var", "tmp", "file.name")
    print(result)
`
	if err := state.DoString(source); err != nil {
		log.Fatal(err.Error())
	}
	// Output:
	// var/tmp/file.name
}

// filepath.glob(string).
func ExampleGlob() {
	state := lua.NewState()
	filepath.Preload(state)
	inspect.Preload(state)
	source := `
    local filepath = require("filepath")
    local inspect = require("inspect")
    local result = filepath.glob("./*/*.lua")
    print(inspect(result, {newline="", indent=""}))
`
	if err := state.DoString(source); err != nil {
		log.Fatal(err.Error())
	}
	// Output:
	// { "test/test_api.lua" }
}
