package grammar

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"
)

type Write struct {
	Expression Expressions
}

type WriteLn Write

func evaluate(a Expressions, context *c.Context) string {
	return strings.Join(lo.Map(
		a, func(e Expression, _ int) string {
			r, _ := e.Evaluate(context)
			return fmt.Sprint(r)
		}), "")
}

func (a Write) Evaluate(context *c.Context) (any, error) {
	result := evaluate(a.Expression, context)
	context.System.OutputText(result)
	return result, nil
}

func (a WriteLn) Evaluate(context *c.Context) (any, error) {
	result := evaluate(a.Expression, context)
	context.System.OutputText(fmt.Sprintf("%s\n", result))
	return result, nil
}

func stringifyList(prefix string, a Expressions) string {
	return fmt.Sprintf("%s %s", prefix,
		strings.Join(
			lo.Map(a,
				func(e Expression, _ int) string {
					return e.String()
				}), " "))
}

func (a Write) String() string {
	return stringifyList("~", a.Expression)
}

func (a WriteLn) String() string {
	return stringifyList("`", a.Expression)
}
