package grammar

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"
)

type NamedExpression c.KeyValue[string, Expression]

type NamedExpressionList []NamedExpression

func (a NamedExpression) String() string {
	if len(a.Key) == 0 {
		return a.Value.String()
	}

	return fmt.Sprintf("%s: %s", a.Key, a.Value.String())
}

func (a NamedExpression) Evaluate(context *c.Context) (any, error) {
	r, err := a.Value.Evaluate(context)
	if err != nil {
		return a, err
	}

	return c.NamedValue{Key: a.Key, Value: r}, nil
}

func (a NamedExpressionList) String() string {
	r := lo.Map(a, func(q NamedExpression, _ int) Expression {
		return Expression(q)
	})

	return strings.Join(stringifyExpressions(r), " ")
}

func (a NamedExpressionList) Evaluate(context *c.Context) (any, error) {
	r := lo.Map(a, func(q NamedExpression, _ int) Expression {
		return Expression(q)
	})

	return Expressions(r).Evaluate(context)
}
