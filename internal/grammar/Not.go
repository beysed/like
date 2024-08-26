package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Not struct {
	Expression Expression
}

func (v Not) Debug() string {
	return fmt.Sprintf("!(%s)", v.Expression.Debug())
}

func (v Not) String() string {
	return fmt.Sprintf("!%s", v.Expression.String())
}

func (a Not) Evaluate(context *c.Context) (any, error) {
	v, err := a.Expression.Evaluate(context)
	if err != nil {
		return nil, err
	}

	str := trim(c.Stringify(v))
	if len(str) > 0 {
		return "", nil
	}
	return "T", nil
}

func MakeNot(s Expression) Expression {
	return Not{Expression: s}
}
