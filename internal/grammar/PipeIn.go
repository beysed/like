package grammar

import (
	"fmt"
	"os"

	c "github.com/beysed/like/internal/grammar/common"
)

type PipeIn Pipe

func (a PipeIn) String() string {
	return fmt.Sprintf("%s < %s", a.to.String(), a.from.String())
}

func (a *PipeIn) From() Ref[Expression] {
	return (*Pipe)(a).From()
}

func (a *PipeIn) To() Ref[Expression] {
	return (*Pipe)(a).To()
}

func (a *PipeIn) Err() Ref[Expression] {
	return (*Pipe)(a).Err()
}

func (a PipeIn) Evaluate(context *c.Context) (any, error) {
	toRef, ok := a.to.(Reference)
	if !ok {
		return a.to, c.MakeError(fmt.Sprintf("%s: should be a reference", a.to.String()), nil)
	}

	from, err := a.from.Evaluate(context)
	if err != nil {
		return a.from, err
	}

	file := c.Stringify(from)
	bytes, err := os.ReadFile(file)
	if err != nil {
		return a.from, c.MakeError(fmt.Sprintf("can not read file: %s", file), err)
	}

	assign := Assign{
		Store: toRef.Expression,
		Value: MakeConstant(string(bytes)),
	}

	return assign.Evaluate(context)
}
