package grammar

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
)

type Expression interface {
	String() string
	Debug() string
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

func EvaluateToString(expression Expression, context *c.Context) (string, error) {
	val, err := expression.Evaluate(context)
	if err != nil {
		return "", err
	}

	if lst, ok := val.([]any); ok {
		result := strings.Builder{}
		for _, l := range lst {
			result.WriteString(fmt.Sprint(l))
		}

		return result.String(), nil
	}

	return fmt.Sprint(val), nil
}
