package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

func Execute(filePath string, context *c.Context, code []byte) (any, error) {
	result, err := Parse(filePath, code, Entrypoint("file"))

	if err != nil {
		//todo: include path? need to test
		return nil, c.MakeError("syntax error", err)
	}

	err = context.PathStack.Push(filePath)
	if err != nil {
		return nil, err
	}

	defer context.PathStack.Pop()

	exprs := result.([]Expression)
	var last any

	for _, expr := range exprs {
		last, err = expr.Evaluate(context)
		_, locals := context.Locals.Peek()
		context.System.OutputText(locals.Output.String())
		context.System.OutputError(locals.Errors.String())
		locals.Reset()

		if err != nil {
			return expr, c.MakeError(fmt.Sprintf("while evaluating: { %s }", expr.String()), err)
		}
	}

	return last, nil
}
