package grammar

import (
	"fmt"
	"os"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"

	"path"
	"strings"
)

func MakeDefaultBuiltIn() c.BuiltIn {
	return c.BuiltIn{
		"exec": func(context *c.Context, args []c.NamedValue) (any, error) {
			if len(args) == 0 {
				return nil, c.MakeError("nothing to exec, too few arguments", nil)
			}

			expressions := Expressions(
				lo.Map(args,
					func(i c.NamedValue, _ int) Expression {
						return MakeConstant([]any{i.Value})
					}))

			invoke := Invoke{
				Expressions: expressions,
			}

			return invoke.Evaluate(context)
		},
		"cwd": func(context *c.Context, args []c.NamedValue) (any, error) {
			if len(args) > 0 {
				return nil, c.MakeError("function does not accept arguments", nil)
			}
			return os.Getwd()
		},
		"file": func(context *c.Context, args []c.NamedValue) (any, error) {
			if len(args) != 1 {
				return nil, c.MakeError("function accepts single argument", nil)
			}

			file, ok := args[0].Value.(string)
			if !ok {
				return nil, c.MakeError("function accepts string argument", nil)
			}

			file, err := context.System.ResolvePath(context, file)
			if err != nil {
				return nil, err
			}
			content, err := os.ReadFile(file)
			return string(content), err
		},
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
