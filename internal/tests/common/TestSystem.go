package common

import (
	"strings"

	"os"

	g "github.com/beysed/like/internal/grammar"
	c "github.com/beysed/like/internal/grammar/common"
)

type TestSystem struct {
	Result strings.Builder
}

func (t *TestSystem) ResolvePath(context *c.Context, filePath string) (string, error) {
	f := File(filePath)
	_, err := os.Stat(f)

	return f, err
}

func (t *TestSystem) OutputText(text string) {
	t.Result.WriteString(text)
}

func (t *TestSystem) OutputError(text string) {
	t.OutputText(text)
}

func MakeContext() (*c.Context, *strings.Builder) {
	system := TestSystem{}
	return g.MakeContext(c.MakeLocals(c.Store{}), g.MakeDefaultBuiltIn(), &system), &system.Result
}
