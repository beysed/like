package grammar

import (
	"fmt"
	"strings"
	"time"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/beysed/shell/execute"
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
			} else if a, ok := e.(Reference); ok {
				res, err := a.Evaluate(context)
				if err != nil {
					return err
				}
				if a, ok := res.([]interface{}); ok {
					for _, u := range a {
						result = append(result, u.(string))
					}
				} else {
					result = append(result, res.(string))
				}
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

	cmd, ok := cmdEval.(string) // todo check
	if !ok {
		return cmdEval, c.MakeError("command is not string", nil)
	}

	args, err := flattern(a.Expressions[1:], context)
	if err != nil {
		return nil, err
	}

	command := execute.MakeCommand(cmd, args...)
	execution, _ := execute.Execute(command)

	close(execution.Stdin)

	for run := true; run; {
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
