package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Pointer struct {
	Expression Expression
}

func (a Pointer) Evaluate(context *c.Context) (any, error) {
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

	return ref.(StoreReference), nil
}

func (a Pointer) String() string {
	r := a.Expression.String()

	return fmt.Sprintf("^%s", r)
}
