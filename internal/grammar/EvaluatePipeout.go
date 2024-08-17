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

func getOutput(ref Ref[Expression], context *c.Context) (string, *Expression, error) {
	a := ref.Get()
	if *a == nil {
		return "", nil, nil
	}

	v, err := (*a).Evaluate(context)
	if err != nil {
		return "", a, err
	}

	return c.Stringify(v), nil, nil
}

func EvaluatePipeout[T PipeOutInstance](a T, context *c.Context, writer OutputWriter) (any, error) {
	_, expr, err := getOutput(a.From(), context)
	if err != nil {
		return expr, err
	}

	_, current := context.Locals.Peek()
	output := current.Output.String()
	errs := current.Errors.String()
	mixed := current.Mixed.String()
	current.Reset()

	aTo, expr, err := getOutput(a.To(), context)
	if err != nil {
		return expr, err
	}

	aErr, expr, err := getOutput(a.Err(), context)
	if err != nil {
		return expr, err
	}

	if aErr == "" {
		err = writer(aTo, mixed)
		if err != nil {
			return aTo, err
		}
	} else {
		err = writer(aTo, output)
		if err != nil {
			return aTo, err
		}
		err = writer(aErr, errs)
		if err != nil {
			return aTo, err
		}
	}

	return a.From(), nil
}
