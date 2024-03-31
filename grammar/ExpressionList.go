package grammar

import (
	"strings"

	"github.com/samber/lo"
)

type ExpressionList []Expression

func (a ExpressionList) String() string {
	return strings.Join(
		lo.Map(a,
			func(v Expression, _ int) string {
				return v.String()
			}), " ")
}

func (a ExpressionList) Evaluate(context *Context) (any, error) {
	var result []any

	for _, v := range a {
		r, err := v.Evaluate(context)
		if err != nil {
			return v, err
		}
		result = append(result, r)
	}

	return result, nil
}