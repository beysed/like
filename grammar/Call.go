package grammar

import "fmt"

type Call struct {
	Store     StoreAccess
	Arguments ExpressionList
}

func (a Call) String() string {
	return fmt.Sprintf("%s(%s)", a.Store.String(), a.Arguments.String())
}

func (a Call) Evaluate(context *Context) (any, error) {
	store, err := a.Store.Evaluate(context)
	if err != nil {
		return store, err
	}

	lambda, ok := store.(Store)[ValueKey].(Lambda)
	if !ok {
		return nil, MakeError(fmt.Sprintf("'%s' is not lambda", a.Store.String()))
	}

	// todo: check len of argument lists

	local := Context{
		Locals:  Store{},
		Globals: context.Globals,
		System:  context.System,
	}

	argc := len(a.Arguments.Expressions)
	for i, v := range lambda.Arguments.Identifiers {
		if i >= argc {
			break
			// todo: warning args mismatch
		}
		arg, err := a.Arguments.Expressions[i].Evaluate(context)
		if err != nil {
			return arg, err
		}

		local.Locals[v] = Store{ValueKey: arg}
	}

	return lambda.Body.Evaluate(&local)
}
