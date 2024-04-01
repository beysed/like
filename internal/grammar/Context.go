package grammar

import (
	"fmt"
	"strings"
)

type Store map[string]any
type BuiltIn map[string]func(context *Context, args []any) (any, error)

type SystemContext struct {
	Buffer strings.Builder
}

func MakeDefaultBuiltIn() BuiltIn {
	return BuiltIn{
		"error": func(context *Context, args []any) (any, error) {
			return nil, MakeError(stringify(args))
		},
		"eval": func(context *Context, args []any) (any, error) {
			var err error

			lines := []string{}
			for _, a := range args {
				var last any

				if t, ok := a.(Expression); ok {
					last, err = t.Evaluate(context)
					if err != nil {
						return t, err
					}
				} else {
					last = a
				}

				lines = append(lines, fmt.Sprint(last))
			}
			text := strings.Join(lines, "\n")
			return Execute(context, []byte(text))
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
