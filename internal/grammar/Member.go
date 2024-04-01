package grammar

type Member struct {
	Identifier string
}

func (a Member) String() string {
	return a.Identifier
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

	if store[a.Identifier] == nil {
		store[a.Identifier] = Store{}
	}

	return store[a.Identifier], nil
}
