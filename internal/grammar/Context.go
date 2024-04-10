package grammar

import (
	c "github.com/beysed/like/internal/grammar/common"
	s "github.com/zeroflucs-given/generics/collections/stack"
)

func MakeContext(locals c.Store, globals c.Store, builtIn c.BuiltIn, system c.System) *c.Context {
	return &c.Context{
		Locals:    locals,
		Globals:   globals,
		BuiltIn:   builtIn,
		System:    system,
		PathStack: s.NewStack[string](128),
	}
}
