package grammar

import (
	"fmt"
	"os"

	c "github.com/beysed/like/internal/grammar/common"
)

type Include struct {
	FileName Expression
}

func (a Include) String() string {
	return fmt.Sprintf("#include %s", a.FileName.String())
}

func (a Include) Evaluate(context *c.Context) (any, error) {
	var v any
	var err error
	var fileName string
	var file []byte

	v, err = a.FileName.Evaluate(context)
	if err != nil {
		return a.FileName, err
	}

	fileName, err = context.System.ResolvePath(context, fmt.Sprint(v))
	if err != nil {
		return fileName, err
	}

	file, err = os.ReadFile(fileName)
	if err != nil {
		return fileName, err
	}

	return Execute(fileName, context, file)
}
