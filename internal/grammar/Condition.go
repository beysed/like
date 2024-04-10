package grammar

import (
	"fmt"
	"strings"

	c "github.com/beysed/like/internal/grammar/common"
)

type Condition struct {
	Condition Expression
	True      Expression
	False     Expression
}

func (a Condition) String() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("? %s", a.Condition.String()))

	if a.True != nil {
		b.WriteString(a.True.String())
	}

	if a.False != nil {
		b.WriteString(fmt.Sprintf("\n%% %s", a.False.String()))
	}

	return b.String()
}

func (a Condition) Evaluate(context *c.Context) (any, error) {
	v, err := a.Condition.Evaluate(context)
	if err != nil {
		return a.Condition, err
	}

	var result bool
	switch t := v.(type) {
	case string:
		result = len(t) > 0
	case c.Store:
		result = len(t) > 0
	case []any:
		result = len(t) > 0
	case any:
		result = t != nil
	case nil:
	default:
		str := trim(stringify(t))
		result = len(str) > 0
	}

	if result {
		if a.True == nil {
			return "", nil
		}

		v, err := a.True.Evaluate(context)
		if err != nil {
			return a.True, err
		}

		return v, nil
	}

	if a.False == nil {
		return "", nil
	}

	v, err = a.False.Evaluate(context)
	if err != nil {
		return a.False, err
	}

	return v, nil
}

func MakeCondition(c Expression, t Expression, f Expression) Expression {
	return Condition{Condition: c, True: t, False: f}
}
