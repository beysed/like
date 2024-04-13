package grammar

import (
	"fmt"
	"os"

	c "github.com/beysed/like/internal/grammar/common"
)

type PipeOut Pipe

func (a *PipeOut) From() Ref[Expression] {
	return (*Pipe)(a).From()
}

func (a *PipeOut) To() Ref[Expression] {
	return (*Pipe)(a).To()
}

func (a PipeOut) String() string {
	return fmt.Sprintf("%s > %s", a.from.String(), a.to.String())
}

func (a PipeOut) Evaluate(context *c.Context) (any, error) {
	return EvaluatePipeout(&a, context, Writer(os.O_WRONLY|os.O_TRUNC|os.O_CREATE))
}
