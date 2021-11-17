package parser

import (
	"regexp"
	"strings"
)

var regexpNolint = regexp.MustCompile(`nolint\:(\S+)`)

// NoLintParse ...
func NoLintParse(query string) []string {
	// -- nolint:test bla-bla
	if !regexpNolint.MatchString(query) {
		return nil
	}
	data := regexpNolint.FindAllStringSubmatch(query, -1)
	result := make([]string, 0)
	for _, m := range data {
		result = append(result, strings.Split(m[1], ",")...)
	}
	return result
}
