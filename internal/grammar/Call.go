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
			l := unwrap(val)
			if la, ok = l.(BindedLambda); !ok {
				return nil, c.MakeError(fmt.Sprintf("'%s' is not lambda", c.Stringify(l)), nil)
			}
		}

		evalFunc = func(context *c.Context, args []c.NamedValue) (any, error) {
			if err != nil {
				return nil, err
			}

			_, current := context.Locals.Peek()

			local := c.MakeLocals(c.Store{})
			local.Store["_i"] = current.Input

			all := c.Store{}
			for k, v := range args {
				all[c.Stringify(k)] = v.Value
			}
			local.Store["_a"] = all

			rest := c.Store{}
			for n := len(la.Lambda.Arguments.Identifiers); n < len(args); n++ {
				rest[c.Stringify(n)] = args[n].Value
			}
			local.Store["_r"] = rest

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

			_, current = context.Locals.Peek()
			current.Output.WriteString(lambdaLocals.Output.String())
			current.Errors.WriteString(lambdaLocals.Errors.String())
			current.Mixed.WriteString(lambdaLocals.Output.String())
			lambdaLocals.Reset()

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
