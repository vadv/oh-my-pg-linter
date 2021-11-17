package rules

// Manager ...
type Manager interface {
	// AddRuleDir add rules from dir.
	AddRuleDir(string) error
	// ListRules is list of rules.
	ListRules() []string
	// Check file with rule.
	Check(file, rule string) (Response, error)
	// Test rule.
	Test(rule string) error
}

// Response ...
type Response interface {
	Passed() bool
	Message() []byte
	Query() *string
}
