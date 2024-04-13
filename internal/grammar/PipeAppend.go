package grammar

import (
	"fmt"
	"os"

	c "github.com/beysed/like/internal/grammar/common"
)

type PipeAppend Pipe

func (a *PipeAppend) From() Ref[Expression] {
	return (*Pipe)(a).From()
}

func (a *PipeAppend) To() Ref[Expression] {
	return (*Pipe)(a).To()
}

func (a PipeAppend) String() string {
	return fmt.Sprintf("%s >> %s", a.from.String(), a.to.String())
}

func (a PipeAppend) Evaluate(context *c.Context) (any, error) {
	return EvaluatePipeout(&a, context, Writer(os.O_WRONLY|os.O_APPEND|os.O_CREATE))
}
