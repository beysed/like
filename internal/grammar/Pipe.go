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
	data, err := a.From.Evaluate(context)
	if err != nil {
		return a.From, err
	}

	locals := c.Store{}
	locals["$_input"] = data
	context.Locals.Push(locals)
	defer context.Locals.Pop()

	return a.To.Evaluate(context)
}

func MakePipe(from Expression, to Expression) Expression {
	return Pipe{From: from, To: to}
}
