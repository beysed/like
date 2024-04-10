package grammar

import (
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"
)

type LiteralList []Literal

func (v LiteralList) String() string {
	return strings.Join(lo.Map(v, func(l Literal, _ int) string { return l.String() }), " ")
}

func (v LiteralList) Evaluate(context *c.Context) (any, error) {
	if len(v) == 1 {
		return v[0].Evaluate(context)
	}

	return lo.Map(v,
		func(l Literal, _ int) any {
			r, _ := l.Evaluate(context)
			return r
		}), nil
}

func MakeLiteralList(s []string) LiteralList {
	return (LiteralList)(lo.Map(s, func(s string, _ int) Literal { return MakeLiteral(s) }))
}
