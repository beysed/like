package grammar

import c "github.com/beysed/like/internal/grammar/common"

type Expression interface {
	String() string
	Evaluate(context *c.Context) (any, error)
}

func Evaluate[T any](expression Expression, context *c.Context) (T, error) {
	var result T
	val, err := expression.Evaluate(context)
	if err != nil {
		return result, err
	}

	if result, ok := val.(T); ok {
		return result, nil
	}

	return result, c.MakeError("incorrect type", nil)
}
