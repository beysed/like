package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Literal struct {
	Value any
}

func (v Literal) Debug() string {
	return v.String()
}

func (v Literal) String() string {
	switch v := v.Value.(type) {
	case fmt.Stringer:
		return v.String()
	case string:
		return v
	default:
		return "unknown"
	}
}

func (v Literal) Evaluate(context *c.Context) (any, error) {
	return v.Value, nil
}

func MakeLiteral(s string) Literal {
	return Literal{Value: s}
}
