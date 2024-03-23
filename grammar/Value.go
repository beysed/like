package grammar

import (
	"fmt"
)

type Value struct {
	V any
}

func (v Value) String() string {
	switch v := v.V.(type) {
	default:
		return "unknown"
	case fmt.Stringer:
		return v.String()
	case string:
		return v
	}
}

func (v Value) Evaluate(system System, globals Context, locals Context) any {
	return v.V
}
