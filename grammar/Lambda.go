package grammar

import "fmt"

type Lambda struct {
	Arguments IdentifierList
	Body      Expression
}

func (a Lambda) String() string {
	return fmt.Sprintf("(%s) %s", a.Arguments.String(), a.Body.String())
}

func (a Lambda) Evaluate(context *Context) (any, error) {
	return a, nil
}
