package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Pipe struct {
	From Expression
	To   Expression
}

func (a Pipe) String() string {
	return fmt.Sprintf("%s | %s", a.From, a.To)
}

func (a Pipe) Evaluate(context *c.Context) (any, error) {
	_, err := a.From.Evaluate(context)
	if err != nil {
		return a.From, err
	}

	_, current := context.Locals.Peek()

	if ref, ok := a.To.(Reference); ok {
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

	return a.To.Evaluate(context)
}

func MakePipe(from Expression, to Expression) Expression {
	return Pipe{From: from, To: to}
}
