package grammar

import c "github.com/beysed/like/internal/grammar/common"

type ExpressionStore c.Store

func (a ExpressionStore) String() string {
	return c.Stringify(c.Store(a))
}

func (a ExpressionStore) Evaluate(context *c.Context) (any, error) {
	res := c.Store{}

	for k, v := range a {
		val, err := v.(Expression).Evaluate(context)
		if err != nil {
			return nil, err
		}

		res[k] = val
	}

	return res, nil
}

func MakeExpressionStore(a NamedExpressionList) (ExpressionStore, error) {
	res := ExpressionStore{}

	for _, v := range a {
		res[v.Key] = v.Value
	}

	return res, nil
}
