package grammar

import (
	"fmt"
)

type Each struct {
	List Expression
	Body Expression
}

func (a Each) String() string {
	return fmt.Sprintf("@ %s %s", a.List.String(), a.Body.String())
}

func (a Each) Evaluate(context *Context) (any, error) {
	v, err := a.List.Evaluate(context)
	if err != nil {
		return a.List, err
	}

	result := []any{}

	var eval = func(l any) (any, error) {
		context.Locals["_"] = Store{ValueKey: l}
		r, err := a.Body.Evaluate(context)
		if err != nil {
			return a.Body, err
		}
		result = append(result, r)

		return r, nil
	}

	if lst, ok := v.([]any); ok {
		for _, l := range lst {
			_, err = eval(l)
			if err != nil {
				return a.Body, err
			}
		}
	} else {
		_, err := eval(v)
		if err != nil {
			return a.Body, err
		}
	}

	return result, nil
}

func MakeEach(l Expression, b Expression) Expression {
	return Each{List: l, Body: b}
}
