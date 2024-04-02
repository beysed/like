package grammar

import c "github.com/beysed/like/internal/grammar/common"

func Execute(context *c.Context, code []byte) (any, error) {
	result, err := Parse("a.like", code, Entrypoint("file"))

	if err != nil {
		return nil, err
	}

	exprs := result.([]Expression)
	var last any
	for _, expr := range exprs {
		last, err = expr.Evaluate(context)

		if err != nil {
			return expr, err
		}
	}

	return last, nil
}
