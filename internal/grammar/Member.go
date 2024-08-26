package grammar

import (
	c "github.com/beysed/like/internal/grammar/common"
	"github.com/zeroflucs-given/generics/collections/stack"
)

type Member struct {
	Identifier string
}

func (a Member) Debug() string {
	return a.String()
}

func (a Member) String() string {
	return a.Identifier
}

func fetchStack[T any](s *stack.Stack[T]) []T {
	fetch := []T{}
	for {
		b, t := s.Pop()
		if !b {
			break
		}

		fetch = append(fetch, t)
	}

	for i := range fetch {
		s.Push(fetch[len(fetch)-1-i])
	}
	return fetch
}

func copyStack[T any](s *stack.Stack[T], size int) *stack.Stack[T] {
	fetch := fetchStack(s)
	c := stack.NewStack[T](size)
	for i := range fetch {
		c.Push(fetch[len(fetch)-1-i])
	}
	return c
}

func findStore(context *c.Context, identifier string) c.Store {
	locals := fetchStack(context.Locals)

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
