package grammar

import (
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
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

type List []any

func (a ExpressionList) Evaluate(context *c.Context) (any, error) {
	var result List

	for _, v := range a {
		r, err := v.Evaluate(context)
		if err != nil {
			return v, err
		}
		result = append(result, r)
	}

	return result, nil
}
