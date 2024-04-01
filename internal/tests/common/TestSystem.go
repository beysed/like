package common

import (
	"fmt"
	"strings"

	g "github.com/beysed/like/internal/grammar"
)

type TestSystem struct {
	Result strings.Builder
}

func (t *TestSystem) ReadFile(filePath string) ([]byte, error) {
	return Read(filePath), nil
}

func (t *TestSystem) Output(text any) {
	t.Result.WriteString(fmt.Sprintf("%s", text))
}

func MakeContext() (g.Context, *strings.Builder) {
	system := TestSystem{}
	store := g.Store{}

	return g.MakeContext(store, store, g.MakeDefaultBuiltIn(), &system), &system.Result
}
