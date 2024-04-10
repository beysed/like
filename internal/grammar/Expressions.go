package grammar

import (
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"
)

type Expressions []Expression

func (a Expressions) String() string {
	return strings.Join(
		lo.Map(a,
			func(e Expression, _ int) string {
				return e.String()
			}), "")
}

func (a Expressions) Evaluate(context *c.Context) (any, error) {
	b := []any{}
	for _, v := range a {
		res, err := v.Evaluate(context)
		if err != nil {
			return v, err
		}

		b = append(b, res)
	}

	return b, nil
}
