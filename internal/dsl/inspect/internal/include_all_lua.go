package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	packageName = "inspect"
	fileName    = "./internal/inspect.lua"
	constName   = "luaInspect"
	templateGo  = `// Package inspect provides inspect.lua for gopher-lua
package %s

// nolint:lll
const %s = "%s"
`
)

func main() {
	out, err := os.Create("lua_const.go")
	if err != nil {
		log.Fatal(err.Error())
	}
	// nolint:errcheck,gosec
	defer out.Close()

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	base46Code := base64.StdEncoding.EncodeToString(data)
	content := fmt.Sprintf(templateGo, packageName, constName, base46Code)
	if _, err := out.WriteString(content); err != nil {
		panic(err)
	}
}
