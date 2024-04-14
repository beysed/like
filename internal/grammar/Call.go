package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Call struct {
	Reference Reference
	Arguments NamedExpressionList
}

func (a Call) String() string {
	return fmt.Sprintf("%s(%s)", a.Reference.String(), a.Arguments.String())
}

func (a Call) Evaluate(context *c.Context) (any, error) {
	var evalFunc func(context *c.Context, args []c.NamedValue) (any, error)

	evalArgs := func() ([]c.NamedValue, any, error) {
		args := []c.NamedValue{}

		for _, v := range a.Arguments {
			arg, err := v.Evaluate(context)
			if err != nil {
				return nil, arg, err
			}

			args = append(args, arg.(c.NamedValue))
		}

		return args, nil, nil
	}

	if store, ok := a.Reference.Expression.(StoreAccess); ok {
		if f, ok := store.Reference.(Literal); ok {
			funcName := f.Value.(string)
			builtInFunc := context.BuiltIn[funcName]
			evalFunc = builtInFunc
		}
	}

	if evalFunc == nil {
		val, err := a.Reference.Evaluate(context)
		if err != nil {
			return a, err
		}

		la, ok := val.(BindedLambda)
		if !ok {
			return nil, c.MakeError(fmt.Sprintf("'%s' is not lambda", a.Reference.String()), nil)
		}

		evalFunc = func(_ *c.Context, args []c.NamedValue) (any, error) {
			if err != nil {
				return nil, err
			}
			if len(args) > len(la.Lambda.Arguments.Identifiers) {
				return nil, c.MakeError("call has extra arguments", nil)
			}

			var named bool
			var unnamed bool
			for _, v := range args {
				if len(v.Key) > 0 {
					named = true
				} else {
					unnamed = true
				}

				if named && unnamed {
					return nil, c.MakeError("can not use named and unnamed args together", nil)
				}
			}

			local := c.MakeLocals(c.Store{})

			if unnamed {
				for i, v := range la.Lambda.Arguments.Identifiers {
					if i >= len(args) {
						break
					}

					local.Store[v] = args[i].Value
				}
			} else {
				m := c.Store{}
				for _, v := range args {
					m[v.Key] = v.Value
				}

				for _, v := range la.Lambda.Arguments.Identifiers {
					if m[v] != nil {
						local.Store[v] = m[v]
					}
				}
			}

			err = la.Context.Locals.Push(local)
			result, err := la.Lambda.Body.Evaluate(la.Context)
			t, lambdaLocals := la.Context.Locals.Pop()
			if !t {
				return nil, c.MakeError("can not pop context", nil)
			}

			_, current := context.Locals.Peek()
			current.Output.WriteString(lambdaLocals.Output.String())
			lambdaLocals.Output.Reset()

			if err != nil {
				return result, err
			}

			return result, err
		}
	}

	args, errArg, err := evalArgs()
	if err != nil {
		return errArg, err
	}

	return evalFunc(context, args)
}
