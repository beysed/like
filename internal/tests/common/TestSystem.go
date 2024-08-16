package common

import (
	"fmt"
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

func (t *TestSystem) Invoke(executable string, args []string, stdin string) (string, string, error) {
	if executable == "fake" {
		return fmt.Sprintf("faked(%s:%s)", stdin, strings.Join(args, ";")), "fake-err", nil
	}

	var a c.CliSystem
	return a.Invoke(executable, args, stdin)
}

func MakeContext() (*c.Context, *strings.Builder) {
	system := TestSystem{}
	locals := c.MakeLocals(c.Store{})
	return g.MakeContext(locals, g.MakeDefaultBuiltIn(), &system), &system.Result
}
