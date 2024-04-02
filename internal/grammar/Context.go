package grammar

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
)

type SystemContext struct {
	Buffer strings.Builder
}

func MakeDefaultBuiltIn() c.BuiltIn {
	return c.BuiltIn{
		"error": func(context *c.Context, args []any) (any, error) {
			return nil, c.MakeError(stringify(args), nil)
		},
		"eval": func(context *c.Context, args []any) (any, error) {
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

func MakeContext(locals c.Store, globals c.Store, builtIn c.BuiltIn, system c.System) c.Context {
	return c.Context{
		Locals:  locals,
		Globals: globals,
		BuiltIn: builtIn,
		System:  system,
	}
}
