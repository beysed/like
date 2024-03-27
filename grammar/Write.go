package grammar

import "fmt"

type Write struct {
	Expression Expression
}

func (a Write) Evaluate(context *Context) (any, error) {
	result, err := a.Expression.Evaluate(context)

	if err != nil {
		return nil, err
	}

	context.System.Output(result)

	return result, nil
}

func (a Write) String() string {
	return fmt.Sprintf("> %s", a.Expression.String())
}
