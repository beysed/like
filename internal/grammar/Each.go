package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Each struct {
	List Expression
	Body Expression
}

func (a Each) String() string {
	return fmt.Sprintf("@ %s %s", a.List.String(), a.Body.String())
}

func (a Each) Evaluate(context *c.Context) (any, error) {
	v, err := a.List.Evaluate(context)
	if err != nil {
		return a.List, err
	}

	result := []any{}

	var eval = func(key string, val any) (any, error) {
		context.Locals["_k"] = fmt.Sprint(key)
		context.Locals["_v"] = fmt.Sprint(val)

		r, err := a.Body.Evaluate(context)
		if err != nil {
			return a.Body, err
		}
		result = append(result, r)

		return r, nil
	}

	//todo map

	if lst, ok := v.([]any); ok {
		for k, l := range lst {
			_, err = eval(fmt.Sprint(k), l)
			if err != nil {
				return a.Body, err
			}
		}
	} else if m, ok := v.(c.Store); ok {
		for k, _ := range m {
			_, err = eval(k, m[k])
			if err != nil {
				return a.Body, err
			}
		}
	} else {
		_, err := eval("0", v)
		if err != nil {
			return a.Body, err
		}
	}

	return result, nil
}

func MakeEach(l Expression, b Expression) Expression {
	return Each{List: l, Body: b}
}
