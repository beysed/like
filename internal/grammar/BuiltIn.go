package grammar

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
)

func MakeDefaultBuiltIn() c.BuiltIn {
	return c.BuiltIn{
		"len": func(context *c.Context, args []any) (any, error) {
			if len(args) != 1 {
				return nil, c.MakeError("'len' accept only single argument", nil)
			}

			r := args[0]
			for {
				if v, ok := r.(Expression); ok {
					r, err := v.Evaluate(context)
					if err != nil {
						return r, err
					}
					continue
				}

				break
			}

			if v, ok := r.([]any); ok {
				return len(v), nil
			}
			return 1, nil
		},

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
			_, folder := context.PathStack.Peek()
			return Execute(folder, context, []byte(text))
		}}
}
