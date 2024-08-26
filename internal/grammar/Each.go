package grammar

import (
	"fmt"

	s "sort"

	c "github.com/beysed/like/internal/grammar/common"
)

type Each struct {
	List Expression
	Body Expression
}

func (a Each) Debug() string {
	return fmt.Sprintf("@(%s %s)", a.List.Debug(), a.Body.Debug())
}

func (a Each) String() string {
	return fmt.Sprintf("@ %s %s", a.List.String(), a.Body.String())
}

func (a Each) Evaluate(context *c.Context) (any, error) {
	v, err := a.List.Evaluate(context)
	if err != nil {
		return a.List, err
	}

	if v == nil {
		return nil, nil
	}

	//result := []any{}
	locals := c.MakeLocals(c.Store{})

	_, current := context.Locals.Peek()
	context.Locals.Push(locals)
	defer context.Locals.Pop()

	var eval = func(key string, val any) (any, error) {
		locals.Store["_k"] = key
		locals.Store["_v"] = val

		r, err := a.Body.Evaluate(context)
		if err != nil {
			return a.Body, err
		}

		return r, nil
	}

	if lst, ok := v.(c.List); ok {
		for k, l := range lst {
			_, err = eval(fmt.Sprint(k), l)
			if err != nil {
				return a.Body, err
			}
		}
	} else if m, ok := v.(c.Store); ok {
		keys := []string{}
		for k := range m {
			keys = append(keys, k)
		}

		s.Strings(keys)
		for _, k := range keys {
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
	current.Output.WriteString(locals.Output.String())
	current.Mixed.WriteString(locals.Output.String())
	locals.Reset()

	return nil, nil
}

func MakeEach(l Expression, b Expression) Expression {
	return Each{List: l, Body: b}
}
