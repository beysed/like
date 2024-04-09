package grammar

import c "github.com/beysed/like/internal/grammar/common"

func Execute(filePath string, context *c.Context, code []byte) (any, error) {
	result, err := Parse("", code, Entrypoint("file"))

	if err != nil {
		return nil, c.MakeError("syntax error", err)
	}

	context.PathStack.Push(filePath)
	defer context.PathStack.Pop()

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
