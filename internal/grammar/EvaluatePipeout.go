package grammar

import (
	"os"

	c "github.com/beysed/like/internal/grammar/common"
)

type OutputWriter func(fname, content string) error

func Writer(mode int) OutputWriter {
	return func(fname string, content string) error {
		file, err := os.OpenFile(fname, mode, 0666)
		defer file.Close()

		if err != nil {
			return err
		}
		_, err = file.Write([]byte(content))
		if err != nil {
			return err
		}

		return nil
	}
}

func EvaluatePipeout[T PipeOutInstance](a T, context *c.Context, writer OutputWriter) (any, error) {
	aFrom := *a.From().Get()
	aTo := *a.To().Get()

	from, err := aFrom.Evaluate(context)
	if err != nil {
		return aFrom, err
	}

	_, current := context.Locals.Peek()
	output := current.Output.String()
	current.Output.Reset()

	to, err := aTo.Evaluate(context)

	if err != nil {
		return aTo, err
	}

	file := c.Stringify(to)
	err = writer(file, output)

	return from, err
}
