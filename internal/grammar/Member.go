package grammar

import c "github.com/beysed/like/internal/grammar/common"

type Member struct {
	Identifier string
}

func (a Member) String() string {
	return a.Identifier
}

func (a Member) Evaluate(context *c.Context) (any, error) {
	var store c.Store
	var stores []c.Store

	if &context.Locals == &context.Globals {
		stores = []c.Store{context.Locals}
	} else {
		stores = []c.Store{context.Locals, context.Globals}
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

	if store[a.Identifier] == nil {
		store[a.Identifier] = c.Store{}
	}

	return store[a.Identifier], nil
}
