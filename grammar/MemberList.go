package grammar

import (
	"strings"

	"github.com/samber/lo"
)

type MemberList []Member

func (a MemberList) String() string {
	return strings.Join(lo.Map(a, func(s Member, _ int) string { return s.String() }), ".")
}

func (a MemberList) Evaluate(context *Context) (any, error) {
	local := Context{
		Globals: context.Globals,
		Locals:  context.Locals,
		System:  context.System}

	var ctx Store
	for _, v := range a {
		res, _ := v.Evaluate(&local)
		ctx = res.(Store)
		local.Locals = ctx
		local.Globals = ctx
	}

	return ctx, nil
}
