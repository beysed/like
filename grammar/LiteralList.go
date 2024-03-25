package grammar

import (
	"strings"

	"github.com/samber/lo"
)

type LiteralList []Literal

func (v LiteralList) String() string {
	return strings.Join(lo.Map(v, func(l Literal, _ int) string { return l.String() }), " ")
}

func LiteralListMake(s []string) LiteralList {
	return (LiteralList)(lo.Map(s, func(s string, _ int) Literal { return LiteralMake(s) }))
}

// func (v Literal) Evaluate(system System, globals Context, locals Context) any {
// 	return v.Value
// }

// func LiteralMake(s string) Literal {
// 	return Literal{Value: s}
// }
