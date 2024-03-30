package grammar

import (
	"fmt"
)

type Reference struct {
	Expression Expression
}

func (a Reference) Evaluate(context *Context) (any, error) {
	var expr Expression
	if e, ok := a.Expression.(Literal); ok {

		i, _ := e.Evaluate(context)
		expr = &StoreAccess{
			Reference: Member{
				Identifier: i.(string),
			}}
	} else {
		expr = a.Expression
	}
	// todo list
	ref, _ := expr.Evaluate(context)

	return ref.(Store)[ValueKey], nil
}

func (a Reference) String() string {
	r := a.Expression.String()

	return fmt.Sprintf("$%s", r)
}
