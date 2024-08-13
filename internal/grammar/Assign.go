package grammar

import (
	c "github.com/beysed/like/internal/grammar/common"
)

type Assign struct {
	Store Expression
	Value Expression
}

func (a Assign) String() string {
	return a.Store.String() + " = " + a.Value.String()
}

func unwrap_single(v any) any {
	for {
		if a, ok := v.([]any); ok {
			if len(a) == 1 {
				v = a[0]
				continue
			}
		}
		return v
	}
}

func (a Assign) Evaluate(context *c.Context) (any, error) {
	store, err := a.Store.Evaluate(context)
	if err != nil {
		return store, err
	}

	if store == nil {
		return a.Store, c.MakeError("assign to nil value", nil)
	}

	v, err := a.Value.Evaluate(context)

	if err != nil {
		return v, err
	}

	v = unwrap_single(v)

	store.(Value).Set(v)

	return v, nil
}
