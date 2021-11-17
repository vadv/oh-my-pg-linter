package parser_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/vadv/oh-my-pg-linter/internal/dsl/parser"
)

var (
	testData = map[string][]string{
		``:                        {},
		`select nolint from test`: {},
		`-- nolint:errcheck select * from query;`: {`errcheck`},
		`-- nolint:errcheck,lol
	select * from query;`: {`errcheck`, `lol`},
		`-- nolint:errcheck
		-- nolint:lol,test
	select * from query;`: {`errcheck`, `lol`, `test`},
	}
)

func TestNoLintParse(t *testing.T) {
	for query, exceptResult := range testData {
		result := parser.NoLintParse(query)
		sort.Strings(result)
		sort.Strings(exceptResult)
		if strings.Join(result, ",") != strings.Join(exceptResult, ",") {
			t.Fatalf("get: %v except: %v", result, exceptResult)
		}
	}
}
