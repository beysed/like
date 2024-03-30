package grammar

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type Write struct {
	Expressions []Expression
}

func (a Write) Evaluate(context *Context) (any, error) {
	result := strings.Join(lo.Map(
		a.Expressions,
		func(e Expression, _ int) string {
			r, _ := e.Evaluate(context)
			return fmt.Sprintf("%s", r)
		}), "")

	context.System.Output(result)

	return result, nil
}

func (a Write) String() string {
	return fmt.Sprintf("` %s",
		strings.Join(
			lo.Map(a.Expressions,
				func(e Expression, _ int) string {
					return e.String()
				}), " "))
}
