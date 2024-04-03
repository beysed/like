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

func evaluate(a Expressions, context *c.Context) (string, error) {
	data := []string{}

	for _, e := range a {
		r, err := e.Evaluate(context)
		if err != nil {
			return "", err
		}

		data = append(data, fmt.Sprint(r))
	}

	return strings.Join(data, ""), nil
}

func (a Write) Evaluate(context *c.Context) (any, error) {
	result, err := evaluate(a.Expression, context)
	if err != nil {
		context.System.OutputError(fmt.Sprintf("%s\n", err.Error()))
		return nil, err
	}

	context.System.OutputText(result)
	return result, nil
}

func (a WriteLn) Evaluate(context *c.Context) (any, error) {
	w := Write{Expression: Expressions{a.Expression, Constant{MakeLiteral("\n")}}}
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

func (a Write) String() string {
	return stringifyList("~", a.Expression)
}

func (a WriteLn) String() string {
	return stringifyList("`", a.Expression)
}
