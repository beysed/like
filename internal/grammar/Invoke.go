package grammar

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	c "github.com/beysed/like/internal/grammar/common"
	p "github.com/beysed/like/internal/grammar/parsers"
	"github.com/beysed/shell/execute"
)

// todo: make something better with these globals
var environ = os.Environ()
var environment, _ = p.GetParser("env").Parse(strings.Join(environ, "\n"))
var shell = environment["LIKE_SH"]

type Invoke struct {
	Expressions Expressions
}

func (a Invoke) String() string {
	return fmt.Sprintf("& %s", a.Expressions.String())
}

func isEmpty(s any) (string, bool) {
	q := s.(string)
	return q, len(strings.TrimSpace(q)) == 0
}

func flattern(exprs []Expression, context *c.Context) ([]string, error) {
	result := []string{}

	var add func(exprs []Expression) error

	add = func(exprs []Expression) error {
		for _, e := range exprs {
			if l, ok := e.(Literal); ok {
				r, e := l.Evaluate(context)
				if e != nil {
					return e
				}
				if s, e := isEmpty(r); e {
					continue
				} else {
					result = append(result, s)
				}
			} else if a, ok := e.(Expressions); ok {
				err := add(a)
				if err != nil {
					return err
				}
			} else if a, ok := e.(ExpressionList); ok {
				err := add(a)
				if err != nil {
					return err
				}
			} else if a, ok := e.(Expression); ok {
				res, err := a.Evaluate(context)
				if err != nil {
					return err
				}

				for _, r := range flat(res) {
					result = append(result, fmt.Sprint(r))
				}
			} else {
				return c.MakeError("unknown element", nil)
			}
		}

		return nil
	}

	err := add(exprs)

	return result, err
}

func (a Invoke) Evaluate(context *c.Context) (any, error) {
	output := strings.Builder{}

	cmdEval, err := a.Expressions[0].Evaluate(context)
	if err != nil {
		return cmdEval, err
	}

	cmd, ok := cmdEval.([]any)
	if !ok {
		return cmdEval, c.MakeError("command is not string", nil)
	}

	args, err := flattern(a.Expressions[1:], context)
	if err != nil {
		return nil, err
	}

	// todo: make lazy, one time
	executable := cmd[0].(string)
	if executable == "$shell" {
		if shell == nil {
			if runtime.GOOS != "windows" {
				executable = "sh"
			} else {
				return nil, c.MakeError("LIKE_SH environmnet variable is not set", nil)
			}
		} else {
			executable = shell.(string)
		}
	}

	command := execute.MakeCommand(executable, args...)
	execution, err := execute.Execute(command)

	if err != nil {
		return nil, err
	}

	close(execution.Stdin)

	run := true
	for run {
		select {
		case out := <-execution.Stdout:
			output.Write(out)
		case <-execution.Exit:
			run = false
		case <-time.After(time.Second * 10):
			execution.Kill()
			run = false
		}
	}

	return output.String(), nil
}
