package grammar

import (
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"
)

type Expressions []Expression

func stringifyExpressions[T ~[]Expression](a T) []string {
	return lo.Map(a,
		func(e Expression, _ int) string {
			return e.String()
		})
}

func flatTop(a []any) []any {
	res := []any{}
	for _, v := range a {
		r, ok := v.([]any)
		if ok {
			for _, f := range r {
				res = append(res, f)
			}
		} else {
			res = append(res, v)
		}
	}
	return res
}

func evaluateExpressions[T ~[]Expression](a T, context *c.Context) ([]any, Expression, error) {
	b := []any{}
	for _, v := range a {
		res, err := v.Evaluate(context)
		if err != nil {
			return nil, v, err
		}

		b = append(b, res)
	}

	return flatTop(b), nil, nil
}

func (a Expressions) String() string {
	return strings.Join(stringifyExpressions(a), "")
}

func (a Expressions) Evaluate(context *c.Context) (any, error) {
	result, errExpr, err := evaluateExpressions(a, context)
	if err != nil {
		return errExpr, err
	}

	return result, nil
}
