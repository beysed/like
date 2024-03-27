package grammar

func Execute(context *Context, code []byte) error {
	result, err := Parse("a.like", code, Entrypoint("file"))

	if err != nil {
		return err
	}

	exprs := arrayify[any](result)

	for _, e := range exprs {
		expr, ok := e.(Expression)

		if !ok {
			continue
		}

		_, err = expr.Evaluate(context)

		if err != nil {
			return err
		}
	}

	return nil
}
