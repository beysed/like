package common

import (
	"fmt"
	"strings"

	g "github.com/beysed/like/internal/grammar"
	c "github.com/beysed/like/internal/grammar/common"
)

type TestSystem struct {
	Result strings.Builder
}

func (t *TestSystem) ReadFile(filePath string) ([]byte, error) {
	return Read(filePath), nil
}

func (t *TestSystem) Output(text any) {
	t.Result.WriteString(fmt.Sprint(text))
}

func MakeContext() (c.Context, *strings.Builder) {
	system := TestSystem{}
	store := c.Store{}

	return g.MakeContext(store, store, g.MakeDefaultBuiltIn(), &system), &system.Result
}
