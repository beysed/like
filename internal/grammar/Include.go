package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Include struct {
	FileName Expression
}

func (a Include) String() string {
	return fmt.Sprintf("#include '%s'", a.FileName.String())
}

func (a Include) Evaluate(context *c.Context) (any, error) {
	var v any
	var err error
	var file []byte

	v, err = a.FileName.Evaluate(context)
	if err != nil {
		return a.FileName, err
	}

	file, err = context.System.ReadFile(fmt.Sprint(v))
	if err != nil {
		return v, err
	}

	return Execute(context, file)
}
