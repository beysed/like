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
	// todo lists
	ref, err := expr.Evaluate(context)
	if err != nil {
		return ref, err
	}

	if ref, ok := ref.(Value); ok {
		return ref.Get(), nil
	}

	return ref, err
}

func (a Reference) String() string {
	r := a.Expression.String()

	return fmt.Sprintf("$%s", r)
}
