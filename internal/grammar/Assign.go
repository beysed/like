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

func (a Assign) Evaluate(context *c.Context) (any, error) {
	store, err := a.Store.Evaluate(context)
	if err != nil {
		return store, err
	}

	v, err := a.Value.Evaluate(context)
	if err != nil {
		return v, err
	}

	store.(Value).Set(v)

	return v, nil
}
