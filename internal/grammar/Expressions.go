package grammar

import (
	"fmt"
	"strings"

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

func (a Expressions) Evaluate(context *Context) (any, error) {
	b := strings.Builder{}
	for _, v := range a {
		res, err := v.Evaluate(context)
		if err != nil {
			return v, err
		}

		b.WriteString(fmt.Sprintf("%s", res))
	}

	return b.String(), nil
}
