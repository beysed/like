package grammar

import (
	"strings"
)

type IdentifierList struct {
	Identifiers []string
}

func (a IdentifierList) String() string {
	return strings.Join(a.Identifiers, " ")
}

func (a IdentifierList) Evaluate(context *Context) (any, error) {
	return a.Identifiers, nil
}
