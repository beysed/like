package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
	p "github.com/beysed/like/internal/grammar/parsers"
)

type ParseData struct {
	Format Expression
	Data   Expression
}

func (v ParseData) String() string {
	return fmt.Sprintf(":< %s %s", v.Format.String(), v.Data.String())
}

func (a ParseData) Evaluate(context *c.Context) (any, error) {
	fmt, err := Evaluate[string](a.Format, context)

	if err != nil {
		return a.Format, c.MakeError("not supported invalid format", nil)
	}

	parser := p.GetParser(fmt)
	if parser == nil {
		return fmt, c.MakeError("no parser for specified format", nil)
	}

	data, err := EvaluateToString(a.Data, context)
	if err != nil {
		return a.Data, err
	}

	parsed, err := parser.Parse(data)
	if err != nil {
		return data, err
	}

	return parsed, nil
}
