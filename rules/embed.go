// Package rules ...
package rules

import (
	"embed"
)

//go:embed *
// Dir with embed files.
var Dir embed.FS
