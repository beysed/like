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
	locals := []*c.Locals{}
	for {
		b, t := s.Pop()
		if !b {
			break
		}

		locals = append(locals, t)
	}

	for i := range locals {
		s.Push(locals[len(locals)-1-i])
	}

	for _, v := range locals {
		if v.Store[identifier] != nil {
			return v.Store
		}
	}

	return nil
}

func (a Member) Evaluate(context *c.Context) (any, error) {
	store := findStore(context, a.Identifier)

	if store == nil {
		_, locals := context.Locals.Peek()
		store = locals.Store
	}

	if store[a.Identifier] == nil {
		return nil, c.MakeError("undefined", nil)
	}

	return store[a.Identifier], nil
}
