package grammar

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type Invoke struct {
	Expressions []Expression
}

func (a *Invoke) Evaluate(context *Context) (any, error) {
	return nil, nil
}

func (a *Invoke) String() string {
	return fmt.Sprintf("& %s",
		strings.Join(
			lo.Map(a.Expressions,
				func(e Expression, _ int) string {
					return e.String()
				}), " "))
}
