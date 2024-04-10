package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
	f "github.com/beysed/like/internal/grammar/formatters"
)

type FormatData struct {
	Format Expression
	Data   Expression
}

func (v FormatData) String() string {
	return fmt.Sprintf(":> %s %s", v.Format.String(), v.Data.String())
}

func (a FormatData) Evaluate(context *c.Context) (any, error) {
	fmt, err := Evaluate[string](a.Format, context)

	if err != nil {
		return a.Format, c.MakeError("not supported invalid format", nil)
	}

	formatter := f.GetFormatter(fmt)
	if formatter == nil {
		return fmt, c.MakeError("no formatter for specified format", nil)
	}

	data, err := Evaluate[c.Store](a.Data, context)
	if err != nil {
		return a.Data, err
	}

	formatted, err := formatter.Format(data)
	if err != nil {
		return data, err
	}

	return formatted, nil
}
