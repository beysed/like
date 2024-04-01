package grammar

import (
	"fmt"
)

type Not struct {
	Expression Expression
}

func (v Not) String() string {
	return fmt.Sprintf("!%s", v.Expression.String())
}

func (a Not) Evaluate(context *Context) (any, error) {
	v, err := a.Expression.Evaluate(context)
	if err != nil {
		return nil, err
	}

	str := trim(stringify(v))
	if len(str) > 0 {
		return "", nil
	}
	return "T", nil
}

func MakeNot(s Expression) Expression {
	return Not{Expression: s}
}
