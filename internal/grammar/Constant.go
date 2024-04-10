package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Constant struct {
	Value any
}

func (a Constant) String() string {
	return fmt.Sprint(a.Value)
}

func (a Constant) Evaluate(context *c.Context) (any, error) {
	return a.Value, nil
}

func MakeConstant(value any) Expression {
	return Constant{Value: value}
}
