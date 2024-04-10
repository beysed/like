package grammar

import c "github.com/beysed/like/internal/grammar/common"

type Member struct {
	Identifier string
}

func (a Member) String() string {
	return a.Identifier
}

func findStore(context *c.Context, identifier string) c.Store {
	s := context.Locals
	stores := []c.Store{}
	for {
		b, t := s.Pop()
		if !b {
			break
		}

		stores = append(stores, t)
	}

	for i, _ := range stores {
		s.Push(stores[len(stores)-1-i])
	}

	for _, v := range stores {
		if v[identifier] != nil {
			return v
		}
	}

	return nil
}

func (a Member) Evaluate(context *c.Context) (any, error) {
	store := findStore(context, a.Identifier)

	if store == nil {
		_, store = context.Locals.Peek()
	}

	if store[a.Identifier] == nil {
		return nil, c.MakeError("undefined", nil)
	}

	return store[a.Identifier], nil
}
