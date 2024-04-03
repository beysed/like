package grammar

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
)

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
			_, folder := context.PathStack.Peek()
			return Execute(folder, context, []byte(text))
		}}
}
