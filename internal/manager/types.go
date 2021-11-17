package manager

import (
	"embed"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Manager ...
type Manager interface {
	// AddRuleDir add rules from dir.
	AddRuleDir(string) error
	// AddEmbed rules.
	AddEmbed(embed.FS) error
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

type box interface {
	ReadFile(string) ([]byte, error)
	Stat(string) error
}

type fsBox struct{}

func (f *fsBox) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Clean(filename))
}

func (f *fsBox) Stat(filename string) (err error) {
	_, err = os.Stat(filepath.Clean(filename))
	return err
}

type embedBox struct {
	embed embed.FS
}

func (f *embedBox) ReadFile(filename string) ([]byte, error) {
	return f.embed.ReadFile(filename)
}

func (f *embedBox) Stat(filename string) error {
	fd, err := f.embed.Open(filename)
	if err != nil {
		return os.ErrNotExist
	}
	if errClose := fd.Close(); errClose != nil {
		return os.ErrNotExist
	}
	return nil
}
