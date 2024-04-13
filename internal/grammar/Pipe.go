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
	_, err := a.from.Evaluate(context)
	if err != nil {
		return a.from, err
	}

	_, current := context.Locals.Peek()

	if ref, ok := a.to.(Reference); ok {
		assign := Assign{
			Store: ref.Expression,
			Value: MakeConstant(current.Output.String())}
		current.Output.Reset()

		return assign.Evaluate(context)
	}

	locals := c.MakeLocals(c.Store{})
	locals.Input = current.Output.String()
	current.Output.Reset()
	context.Locals.Push(locals)

	defer func() {
		current.Output.WriteString(locals.Output.String())
		context.Locals.Pop()
	}()

	return a.to.Evaluate(context)
}
