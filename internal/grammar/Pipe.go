package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Pipe struct {
	from Expression
	to   Expression
}

func (a *Pipe) From() Ref[Expression] {
	return MakeRef(&a.from)
}

func (a *Pipe) To() Ref[Expression] {
	return MakeRef(&a.to)
}

func (a Pipe) String() string {
	return fmt.Sprintf("%s | %s", a.from, a.to)
}

func (a Pipe) Evaluate(context *c.Context) (any, error) {
	res_from, err := a.from.Evaluate(context)
	if err != nil {
		return a.from, err
	}

	var input string
	_, current := context.Locals.Peek()

	_, isRef := a.from.(Reference)
	_, isLit := a.from.(Literal)
	_, isConst := a.from.(Constant)
	_, isString := a.from.(ParsedString)
	_, isExpressions := a.from.(Expressions)

	if isRef || isLit || isConst || isString || isExpressions {
		input = c.Stringify(res_from)
	} else {
		input = current.Output.String()
		current.Output.Reset()
	}

	if ref, ok := a.to.(Reference); ok {
		assign := Assign{
			Store: ref.Expression,
			Value: MakeConstant(input)}

		return assign.Evaluate(context)
	}

	locals := c.MakeLocals(c.Store{})
	locals.Input = input
	context.Locals.Push(locals)

	res, err := a.to.Evaluate(context)

	current.Output.WriteString(locals.Output.String())
	context.Locals.Pop()

	return res, err
}
