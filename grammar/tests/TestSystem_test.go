package tests

import (
	"strings"
)

type TestSystem struct {
	Result strings.Builder
}

func (t TestSystem) Output(text string) {
	t.Result.WriteString(text)
}

func (t TestSystem) Invoke(command string, args ...[]string) (string, error) {
	return "", nil
}
