package rules_test

import (
	"testing"

	"github.com/vadv/oh-my-pg-linter/internal/rules"
)

func TestCheck(t *testing.T) {
	m, err := rules.New("./tests/rules")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := m.Check("", "unknown"); err == nil {
		t.Fatal("must be error")
	}
	for _, r := range m.ListRules() {
		res, err := m.Check("./tests/migrations/test.sql", r)
		if err != nil {
			t.Fatal(err)
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
			if string(res.Message()) != "error" {
				t.Fatal(res.Message())
			}
		}
	}
}

func TestRuleTest(t *testing.T) {
	m, err := rules.New("./tests/rules")
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Test(`rule_3`); err != nil {
		t.Fatal(err)
	}
}
