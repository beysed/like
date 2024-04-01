package grammar

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type Block []Expression

func (a Block) String() string {
	body := strings.Join(
		lo.Map(a,
			func(v Expression, _ int) string {
				return v.String()
			}), "\n")

	return fmt.Sprintf("{\n%s\n}", body)
}

func (a Block) Evaluate(context *Context) (any, error) {
	var last any
	for _, v := range a {
		r, err := v.Evaluate(context)
		if err != nil {
			return v, err
		}
		last = r
	}

	return last, nil
}
