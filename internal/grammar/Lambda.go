package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Lambda struct {
	Arguments IdentifierList
	Body      Expression
}

func (a Lambda) String() string {
	return fmt.Sprintf("(%s) %s", a.Arguments.String(), a.Body.String())
}

func (a Lambda) Evaluate(context *c.Context) (any, error) {
	return a, nil
}
