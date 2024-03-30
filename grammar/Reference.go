package grammar

import (
	"fmt"
	"strings"
)

type Reference struct {
	Expression Expression
}

func (a Reference) Evaluate(context *Context) (any, error) {
	var expr Expression
	if e, ok := a.Expression.(LiteralList); ok {
		// todo list
		i, _ := e.Evaluate(context)
		expr = &MemberList{Member{
			Identifier: i.(string),
		}}
	} else {
		expr = a.Expression
	}

	ref, _ := expr.Evaluate(context)

	return ref.(Store)["value"], nil
}

func (a Reference) String() string {
	r := a.Expression.String()
	var format string
	if strings.Contains(r, ".") {
		format = "$(%s)"
	} else {
		format = "$%s"
	}

	return fmt.Sprintf(format, r)
}
