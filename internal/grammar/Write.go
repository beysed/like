package grammar

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"
)

type Write struct {
	Expression Expressions
	Error      bool
}

type WriteLn Write

func flat(a any) []any {
	result := []any{}

	var r func(any)
	r = func(t any) {
		if a, ok := t.([]any); ok {
			for _, z := range a {
				r(z)
			}
			return
		}

		if a, ok := t.(c.List); ok {
			for _, z := range a {
				r(z)
			}
			return
		}

		result = append(result, t)
	}

	r(a)

	return result
}

func halfflat(a any) []any {
	result := []any{}

	var r func(any)
	r = func(t any) {
		if a, ok := t.([]any); ok {
			for _, z := range a {
				r(z)
			}
			return
		}

		if a, ok := t.([]any); ok {
			for _, z := range a {
				r(z)
			}
			return
		}

		result = append(result, t)
	}

	r(a)

	return result
}

func evaluate(a Expressions, context *c.Context) (string, error) {
	data := []string{}

	for _, e := range a {
		r, err := e.Evaluate(context)
		if err != nil {
			return "", err
		}

		for _, s := range halfflat(r) {
			data = append(data, c.Stringify(s))
		}
	}

	return strings.Join(data, ""), nil
}

func (a Write) Evaluate(context *c.Context) (any, error) {
	result, err := evaluate(a.Expression, context)
	if err != nil {
		context.System.OutputError(fmt.Sprintf("%s\n", err.Error()))
		return nil, err
	}

	_, locals := context.Locals.Peek()

	var output *strings.Builder
	if a.Error {
		output = &locals.Errors
	} else {
		output = &locals.Output
	}

	output.WriteString(result)
	locals.Mixed.WriteString(result)

	return nil, nil
}

func (a WriteLn) Evaluate(context *c.Context) (any, error) {
	w := Write{
		Expression: Expressions{
			a.Expression, Constant{MakeLiteral("\n")}}, Error: a.Error}

	return w.Evaluate(context)
}

func stringifyList(prefix string, a Expressions) string {
	return fmt.Sprintf("%s %s", prefix,
		strings.Join(
			lo.Map(a,
				func(e Expression, _ int) string {
					return e.String()
				}), " "))
}

func debugifyList(prefix string, a Expressions) string {
	return fmt.Sprintf("%s(%s)", prefix,
		strings.Join(
			lo.Map(a,
				func(e Expression, _ int) string {
					return e.Debug()
				}), " "))
}

func getWrite(op string, err bool) string {
	var suf string
	if err {
		suf = "*"
	}

	return fmt.Sprintf("%s%s", op, suf)
}

func (a Write) Debug() string {
	return debugifyList(getWrite("~", a.Error), a.Expression)
}

func (a WriteLn) Debug() string {
	return debugifyList(getWrite("`", a.Error), a.Expression)
}

func (a Write) String() string {
	return stringifyList(getWrite("~", a.Error), a.Expression)
}

func (a WriteLn) String() string {
	return stringifyList(getWrite("`", a.Error), a.Expression)
}
