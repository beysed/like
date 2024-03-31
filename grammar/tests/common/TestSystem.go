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

func (t *TestSystem) Invoke(command string, args ...[]string) (any, error) {
	return nil, nil
}

func MakeContext() (g.Context, *strings.Builder) {
	system := TestSystem{}
	store := g.Store{}

	return g.Context{
		System:  &system,
		Locals:  store,
		Globals: store,
	}, &system.Result
}
