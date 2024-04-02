package grammar

import c "github.com/beysed/like/internal/grammar/common"

func Execute(filePath string, context *c.Context, code []byte) (any, error) {
	result, err := Parse(filePath, code, Entrypoint("file"))
	context.PathStack.Push(filePath)
	defer context.PathStack.Pop()

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
