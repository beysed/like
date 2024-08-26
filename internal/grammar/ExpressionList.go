package grammar

import (
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
)

type ExpressionList []Expression

func (a ExpressionList) Debug() string {
	return strings.Join(debugifyExpressions(a), " ")
}

func (a ExpressionList) String() string {
	return strings.Join(stringifyExpressions(a), " ")
}

func (a ExpressionList) Evaluate(context *c.Context) (any, error) {
	res, err := Expressions(a).Evaluate(context)
	if err != nil {
		return res, err
	}

	return c.List(res.([]any)), nil
}
