package grammar

import (
	c "github.com/beysed/like/internal/grammar/common"
)

func MakeDefaultContextFor(system c.System) *c.Context {
	return MakeContext(c.Store{}, c.Store{}, MakeDefaultBuiltIn(), system)
}

func MakeContext(locals c.Store, globals c.Store, builtIn c.BuiltIn, system c.System) *c.Context {
	return &c.Context{
		Locals:  locals,
		Globals: globals,
		BuiltIn: builtIn,
		System:  system,
	}
}
