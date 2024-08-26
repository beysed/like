package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Reference struct {
	Expression Expression
}

func MakeStoreAccess(member string) *StoreAccess {
	return &StoreAccess{
		Reference: Member{
			Identifier: member}}
}

func (a Reference) Evaluate(context *c.Context) (any, error) {
	var expr Expression

	if e, ok := a.Expression.(Literal); ok {
		i, err := e.Evaluate(context)
		if err != nil {
			return e, err
		}

		expr = MakeStoreAccess(i.(string))
	} else if e, ok := a.Expression.(Reference); ok {
		i, err := e.Evaluate(context)
		if err != nil {
			return e, err
		}

		if s, ok := i.(string); ok {
			expr = MakeStoreAccess(s)
		} else {
			return i, c.MakeError("reference: don't know what to do", nil)
		}
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

func (a Reference) Debug() string {
	return a.String()
}

func (a Reference) String() string {
	r := a.Expression.String()

	return fmt.Sprintf("$%s", r)
}
