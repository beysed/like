package grammar

import (
	"fmt"
)

type Equal struct {
	Left  Expression
	Right Expression
}

func (v Equal) String() string {
	return fmt.Sprintf("%s == %s", v.Left.String(), v.Right.String())
}

func (a Equal) Evaluate(context *Context) (any, error) {
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

	if ls == rs {
		return "T", nil
	}

	return "", nil
}

func MakeEqual(l Expression, r Expression) Expression {
	return Equal{Left: l, Right: r}
}
