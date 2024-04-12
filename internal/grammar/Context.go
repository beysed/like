package grammar

import (
	c "github.com/beysed/like/internal/grammar/common"
	s "github.com/zeroflucs-given/generics/collections/stack"
)

func MakeContext(locals *c.Locals, builtIn c.BuiltIn, system c.System) *c.Context {
	l := s.NewStack[*c.Locals](128)
	l.Push(locals)
	return &c.Context{
		Locals:    l,
		BuiltIn:   builtIn,
		System:    system,
		PathStack: s.NewStack[string](128),
	}
}
