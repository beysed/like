package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Call struct {
	Store     StoreAccess
	Arguments ExpressionList
}

func (a Call) String() string {
	return fmt.Sprintf("%s(%s)", a.Store.String(), a.Arguments.String())
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

	if f, ok := a.Store.Reference.(Literal); ok {
		funcName := f.Value.(string)
		builtInFunc := context.BuiltIn[funcName]
		evalFunc = builtInFunc
	}

	if evalFunc == nil {
		store, err := a.Store.Evaluate(context)
		if err != nil {
			return store, err
		}

		lambda, ok := store.(Value).Get().(Lambda)
		if !ok {
			return nil, c.MakeError(fmt.Sprintf("'%s' is not lambda", a.Store.String()), nil)
		}

		// todo: check len of argument lists

		evalFunc = func(_ *c.Context, args []any) (any, error) {
			local := MakeContext(c.Store{}, context.Globals, context.BuiltIn, context.System)

			argc := len(args)
			for i, v := range lambda.Arguments.Identifiers {
				if i >= argc {
					return nil, c.MakeError("lambda arguments mismatch", nil)
				}

				local.Locals[v] = args[i]
			}

			return lambda.Body.Evaluate(local)
		}
	}

	args, err := evalArgs()
	if err != nil {
		return nil, err
	}

	return evalFunc(context, args)
}
