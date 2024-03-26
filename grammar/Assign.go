package grammar

type Assign struct {
	Identifier string
	Value      Expression
}

func (a Assign) String() string {
	return a.Identifier + " = " + a.Value.String()
}

func (a Assign) Evaluate(context *Context) (any, error) {
	var v any
	var err error

	if a.Value != nil {
		v, err = a.Value.Evaluate(context)
	}

	if err == nil {
		context.Locals[a.Identifier] = v
	}

	return v, err
}
