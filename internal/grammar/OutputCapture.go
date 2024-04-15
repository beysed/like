package grammar

import (
	c "github.com/beysed/like/internal/grammar/common"
)

type OutputCapture Expressions

func (a OutputCapture) String() string {
	return Expressions(a).String()
}

func (a OutputCapture) Evaluate(context *c.Context) (any, error) {
	locals := c.MakeLocals(c.Store{})
	context.Locals.Push(locals)

	for _, v := range a {
		res, err := v.Evaluate(context)
		if err != nil {
			return v, err
		}

		locals.Output.Write([]byte(c.Stringify(unwrap(res))))
	}
	context.Locals.Pop()

	return locals.Output.String(), nil
}
