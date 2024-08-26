package tests

import (
	"os"

	g "github.com/beysed/like/internal/grammar"
	. "github.com/beysed/like/internal/tests/common"
)

func Evaluate(code string) (*TestSystem, any, error) {
	context, system := MakeContext()
	wd, _ := os.Getwd()

	_, locals := context.Locals.Peek()
	locals.Store["_shell"] = GetShell()
	result, err := g.Execute(wd, context, []byte(code))

	return system, result, err
}
