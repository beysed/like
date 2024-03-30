package grammar

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type Member struct {
	Identifier string
	Indexes    []Expression
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

func (a Member) Evaluate(context *Context) (any, error) {
	var store Store
	var stores []Store

	if &context.Locals == &context.Globals {
		stores = []Store{context.Locals}
	} else {
		stores = []Store{context.Locals, context.Globals}
	}

	for _, v := range stores {
		if v[a.Identifier] != nil {
			store = v
			break
		}
	}

	if store == nil {
		store = context.Locals
	}

	// todo: indexes
	//isArray := len(a.Indexes) > 0
	if store[a.Identifier] == nil {
		store[a.Identifier] = Store{}
	}

	return store[a.Identifier], nil
}
