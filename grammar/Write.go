package grammar

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type Write struct {
	Expression Expressions
}

type WriteLn Write

func evaluate(a Expressions, context *Context) string {
	return strings.Join(lo.Map(
		a, func(e Expression, _ int) string {
			r, _ := e.Evaluate(context)
			return fmt.Sprintf("%s", r)
		}), "")
}

func (a Write) Evaluate(context *Context) (any, error) {
	result := evaluate(a.Expression, context)
	context.System.Output(result)
	return result, nil
}

func (a WriteLn) Evaluate(context *Context) (any, error) {
	result := evaluate(a.Expression, context)
	context.System.Output(fmt.Sprintf("%s\n", result))
	return result, nil
}

func stringify(prefix string, a Expressions) string {
	return fmt.Sprintf("%s %s", prefix,
		strings.Join(
			lo.Map(a,
				func(e Expression, _ int) string {
					return e.String()
				}), " "))
}

func (a Write) String() string {
	return stringify("~", a.Expression)
}

func (a WriteLn) String() string {
	return stringify("`", a.Expression)
}
