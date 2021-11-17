package manager_test

import (
	"strings"
	"testing"

	"github.com/vadv/oh-my-pg-linter/internal/manager"
)

func TestCheck(t *testing.T) {
	m := manager.New()
	if err := m.AddRuleDir("./tests/rules"); err != nil {
		t.Fatal(err)
	}
	if _, err := m.Check("", "unknown"); err == nil {
		t.Fatal("must be error")
	}
	for _, r := range m.ListRules() {
		res, err := m.Check("./tests/migrations/test.sql", r)
		if err != nil {
			t.Fatalf("test rule %s: %s", r, err)
		}
		switch r {
		case "rule_1":
			if !res.Passed() {
				t.Fatal("must passed")
			}
		case "rule_2":
			if res.Passed() {
				t.Fatal("must not passed")
			}
			if string(res.Message()) != "rule_2 test\n" {
				t.Fatalf("%s", res.Message())
			}
		}
	}
}

func TestRuleTest(t *testing.T) {
	m := manager.New()
	if err := m.AddRuleDir("./tests/rules"); err != nil {
		t.Fatal(err)
	}
	if err := m.Test(`rule_3`); err != nil {
		t.Fatal(err)
	}
}

func TestRule2NoLint(t *testing.T) {
	m := manager.New()
	if err := m.AddRuleDir("./tests/rules"); err != nil {
		t.Fatal(err)
	}
	r, errCheck := m.Check("./tests/migrations/test_no_lint.sql", "rule_2")
	if errCheck != nil {
		t.Fatal(errCheck)
	}
	if r.Passed() {
		t.Fatal(r.Passed())
	}
	if q := *r.Query(); !strings.Contains(q, `email_must_lint_idx`) {
		t.Fatal(q)
	}
}
