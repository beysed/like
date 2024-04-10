package grammar

import (
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
)

type IdentifierList struct {
	Identifiers []string
}

func (a IdentifierList) String() string {
	return strings.Join(a.Identifiers, " ")
}

func (a IdentifierList) Evaluate(context *c.Context) (any, error) {
	return a.Identifiers, nil
}
