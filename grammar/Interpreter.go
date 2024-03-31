package grammar

func Execute(context *Context, code []byte) error {
	result, err := Parse("a.like", code, Entrypoint("file"))

	if err != nil {
		return err
	}

	exprs := result.([]Expression)

	for _, expr := range exprs {
		_, err = expr.Evaluate(context)

		if err != nil {
			return err
		}
	}

	return nil
}
