package grammar

import (
	"fmt"
	"strings"
)

type Store map[string]any
type BuiltIn map[string]func([]any) (any, error)

type SystemContext struct {
	Buffer strings.Builder
}

func MakeDefaultBuiltIn() BuiltIn {
	return BuiltIn{
		"error": func(a []any) (any, error) {
			return nil, MakeError(stringify(a))
		}}
}

func MakeSystemContext() SystemContext {
	return SystemContext{}
}

func (c *SystemContext) Output(text string) {
	fmt.Print(text)
}

type Context struct {
	Locals  Store
	Globals Store
	BuiltIn BuiltIn
	System  System
}

func MakeContext(locals Store, globals Store, builtIn BuiltIn, system System) Context {
	return Context{
		Locals:  locals,
		Globals: globals,
		BuiltIn: builtIn,
		System:  system,
	}
}
