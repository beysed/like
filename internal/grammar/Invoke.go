package grammar

import (
	"fmt"
	"runtime"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"
)

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

				f := flat(res)
				for _, r := range f {
					s := fmt.Sprint(r)
					if len(strings.TrimSpace(s)) == 0 {
						continue
					}

					result = append(result, s)
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
	cmdEval, err := a.Expressions[0].Evaluate(context)
	if err != nil {
		return a.Expressions[0], err
	}

	cmd, ok := cmdEval.([]any)
	if !ok {
		return cmdEval, c.MakeError("command is not string", nil)
	}
	cmd = flat(cmd)

	args, err := flattern(a.Expressions[1:], context)
	if err != nil {
		return nil, err
	}

	args = append(
		lo.Map(cmd[1:],
			func(v any, _ int) string {
				return c.Stringify(v)
			}), args...)

	// todo: make lazy, one time
	var executable string
	if cmd[0] != nil {
		executable = cmd[0].(string)
	}

	if executable == "" {
		if runtime.GOOS != "windows" {
			executable = "sh"
		} else {
			return nil, c.MakeError("LIKE_SH environment variable is not set", nil)
		}
	}

	if err != nil {
		return nil, err
	}

	_, locals := context.Locals.Peek()
	input := locals.Input

	stdout, _, err := context.System.Invoke(executable, args, input)
	locals.Output.WriteString(stdout)

	return err, nil
}
