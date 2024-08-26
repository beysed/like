package common

import (
	"fmt"
	"strings"

	"os"

	g "github.com/beysed/like/internal/grammar"
	c "github.com/beysed/like/internal/grammar/common"
)

type TestSystem struct {
	Stdout strings.Builder
	Stderr strings.Builder
	Stdmix strings.Builder
}

func (t *TestSystem) ResolvePath(context *c.Context, filePath string) (string, error) {
	f := File(filePath)
	_, err := os.Stat(f)

	return f, err
}

func (t *TestSystem) OutputText(text string) {
	t.Stdout.WriteString(text)
	t.Stdmix.WriteString(text)
}

func (t *TestSystem) OutputError(text string) {
	t.Stderr.WriteString(text)
	t.Stdmix.WriteString(text)
}

func (t *TestSystem) Invoke(executable string, args []string, stdin string) (string, string, string, error) {
	if executable == "fake" {
		stdout := fmt.Sprintf("faked(stdin:%s; args:%s)", stdin, strings.Join(args, ", "))
		stderr := "fake-err"
		return stdout, stderr, stdout + stderr, nil
	}

	var a c.CliSystem
	return a.Invoke(executable, args, stdin)
}

func MakeContext() (*c.Context, *TestSystem) {
	system := TestSystem{}
	locals := c.MakeLocals(c.Store{})
	return g.MakeContext(locals, g.MakeDefaultBuiltIn(), &system), &system
}
