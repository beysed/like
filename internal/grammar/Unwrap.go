package grammar

import (
	c "github.com/beysed/like/internal/grammar/common"
)

type Unwrap struct {
	Expression Expression
}

func (v Unwrap) Debug() string {
	return v.Expression.Debug()
}

func (v Unwrap) String() string {
	return v.Expression.String()
}

func (v Unwrap) Evaluate(context *c.Context) (any, error) {
	val, err := v.Expression.Evaluate(context)
	if err != nil {
		return v, err
	}

	return unwrap_single(val), nil
}

func MakeUnwrap(s Expression) Expression {
	return Unwrap{Expression: s}
}
