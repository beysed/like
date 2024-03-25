package grammar

import (
	"fmt"
	"strings"

	. "like/expressions"

	"github.com/samber/lo"
)

type Member struct {
	Identifier string
	Indexes    []Expression
}

type MemberList []Member

func (a MemberList) String() string {
	return strings.Join(lo.Map(a, func(s Member, _ int) string { return s.String() }), ".")
}

func (a Member) String() string {

	if a.Indexes == nil || len(a.Indexes) == 0 {
		return a.Identifier
	}

	return fmt.Sprintf("%s%s", a.Identifier, strings.Join(lo.Map(a.Indexes,
		func(e Expression, _ int) string {
			return fmt.Sprintf("[%s]", e.String())
		}), ""))
}

func (a Member) Evaluate(system System, context *Context) (any, error) {
	return nil, nil
}

func (a MemberList) Evaluate(system System, context *Context) (any, error) {
	return nil, nil
}
