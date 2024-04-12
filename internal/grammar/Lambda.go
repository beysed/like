package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Lambda struct {
	Arguments IdentifierList
	Body      Expression
}

type BindedLambda struct {
	Lambda  *Lambda
	Context *c.Context
}

func (a BindedLambda) Evaluate(context *c.Context) (any, error) {
	return nil, c.MakeError("can not be evaluated directly", nil)
}

func (a BindedLambda) String() string {
	return a.Lambda.String()
}

func (a Lambda) String() string {
	return fmt.Sprintf("(%s) %s", a.Arguments.String(), a.Body.String())
}

func copyContext(context *c.Context) *c.Context {
	c := c.Context{
		Locals:    copyStack(context.Locals, 128),
		BuiltIn:   context.BuiltIn,
		System:    context.System,
		PathStack: copyStack(context.PathStack, 128),
	}

	return &c
}

func (a Lambda) Evaluate(context *c.Context) (any, error) {
	return BindedLambda{
		Lambda:  &a,
		Context: copyContext(context),
	}, nil
}
