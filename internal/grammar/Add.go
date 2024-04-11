package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Add struct {
	Left  Expression
	Right Expression
}

func (v Add) String() string {
	return fmt.Sprintf("%s + %s", v.Left.String(), v.Right.String())
}

func (a Add) Evaluate(context *c.Context) (any, error) {
	l, err := a.Left.Evaluate(context)
	if err != nil {
		return a.Left, err
	}

	r, err := a.Right.Evaluate(context)
	if err != nil {
		return a.Right, err
	}

	ls := stringify(l)
	rs := stringify(r)

	return fmt.Sprintf("%s%s", ls, rs), nil
}

func MakeAdd(l Expression, r Expression) Expression {
	return Add{Left: l, Right: r}
}
