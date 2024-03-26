package grammar

import (
	"fmt"
)

type Literal struct {
	Value any
}

func (v Literal) String() string {
	switch v := v.Value.(type) {
	default:
		return "unknown"
	case fmt.Stringer:
		return v.String()
	case string:
		return v
	}
}

func (v Literal) Evaluate(context *Context) (any, error) {
	return v.Value, nil
}

func LiteralMake(s string) Literal {
	return Literal{Value: s}
}
