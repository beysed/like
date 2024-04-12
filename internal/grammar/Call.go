package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Call struct {
	Reference Reference
	Arguments ExpressionList
}

func (a Call) String() string {
	return fmt.Sprintf("%s(%s)", a.Reference.String(), a.Arguments.String())
}

func (a Call) Evaluate(context *c.Context) (any, error) {
	var evalFunc func(context *c.Context, args []any) (any, error)

	evalArgs := func() ([]any, error) {
		args := []any{}

		for _, v := range a.Arguments {
			arg, err := v.Evaluate(context)
			if err != nil {
				return nil, err
			}

			args = append(args, arg)
		}

		return args, nil
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

		// todo: check len of argument lists

		evalFunc = func(_ *c.Context, args []any) (any, error) {
			local := c.MakeLocals(c.Store{})

			argc := len(args)
			for i, v := range la.Lambda.Arguments.Identifiers {
				if i >= argc {
					return nil, c.MakeError("lambda arguments mismatch", nil)
				}

				local.Store[v] = args[i]
			}

			err = la.Context.Locals.Push(local)
			if err != nil {
				return nil, err
			}

			result, err := la.Lambda.Body.Evaluate(la.Context)
			t, lambdaLocals := la.Context.Locals.Pop()
			_, current := context.Locals.Peek()
			current.Output.WriteString(lambdaLocals.Output.String())

			if !t {
				return nil, c.MakeError("can not pop context", nil)
			}

			if err != nil {
				return result, err
			}

			return result, err
		}
	}

	args, err := evalArgs()
	if err != nil {
		return nil, err
	}

	return evalFunc(context, args)
}
