package common

import (
	"fmt"
	g "like/grammar"
	"strings"
)

type TestSystem struct {
	Result strings.Builder
}

func (t *TestSystem) Output(text any) {
	t.Result.WriteString(fmt.Sprintf("%s", text))
}

func (t *TestSystem) Invoke(command string, args ...[]string) (g.InvokeResult, error) {
	return g.InvokeResult{Stdout: "{ \"a\" : 1 }"}, nil
}
