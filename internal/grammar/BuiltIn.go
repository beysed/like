package grammar

import (
	"fmt"
	"path"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"
)

func MakeDefaultBuiltIn() c.BuiltIn {
	return c.BuiltIn{
		"joinPath": func(context *c.Context, args []c.NamedValue) (any, error) {
			return path.Join(
				lo.Map(args,
					func(a c.NamedValue, _ int) string {
						return c.Stringify(a.Value)
					})...), nil
		},
		"resolvePath": func(context *c.Context, args []c.NamedValue) (any, error) {
			if len(args) != 1 {
				return nil, c.MakeError("'resolvePath' accept only single argument", nil)
			}
			return context.System.ResolvePath(context, c.Stringify(args[0]))
		},
		"len": func(context *c.Context, args []c.NamedValue) (any, error) {
			if len(args) != 1 {
				return nil, c.MakeError("'len' accept only single argument", nil)
			}

			r := args[0]
			for {
				if v, ok := r.Value.(Expression); ok {
					r, err := v.Evaluate(context)
					if err != nil {
						return r, err
					}
					continue
				}

				break
			}

			if v, ok := r.Value.(c.List); ok {
				return len(v), nil
			}

			if v, ok := r.Value.(c.Store); ok {
				return len(v), nil
			}

			return 1, nil
		},
		"error": func(context *c.Context, args []c.NamedValue) (any, error) {
			err := c.Stringify(args)
			return nil, c.MakeError(err, nil)
		},
		"eval": func(context *c.Context, args []c.NamedValue) (any, error) {
			var err error

			lines := []string{}
			for _, a := range args {
				var last any

				if t, ok := a.Value.(Expression); ok {
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

			// to avoid duplicated output
			context.Locals.Push(c.MakeLocals(c.Store{}))
			defer context.Locals.Pop()

			return Execute(folder, context, []byte(text))
		}}
}
