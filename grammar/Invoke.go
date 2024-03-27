package grammar

import (
	"fmt"
)

type Invoke struct {
	Expression Expression
}

func (a *Invoke) Evaluate(context *Context) (any, error) {
	return nil, nil
}

func (a *Invoke) String() string {
	return fmt.Sprintf("& %s", a.Expression.String())
}
