package grammar

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
	"github.com/samber/lo"
)

type Block []Expression

func (a Block) Debug() string {
	body := strings.Join(
		lo.Map(a,
			func(v Expression, _ int) string {
				return v.Debug()
			}), "\n")

	return fmt.Sprintf("{\n%s\n}", body)
}

func (a Block) String() string {
	body := strings.Join(
		lo.Map(a,
			func(v Expression, _ int) string {
				return v.String()
			}), "\n")

	return fmt.Sprintf("{\n%s\n}", body)
}

func (a Block) Evaluate(context *c.Context) (any, error) {
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
