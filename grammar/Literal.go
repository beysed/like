package grammar

import (
	"fmt"
)

type Literal struct {
	Value any
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

func (v Literal) Evaluate(context *Context) (any, error) {
	return v.Value, nil
}

func MakeLiteral(s string) Literal {
	return Literal{Value: s}
}
