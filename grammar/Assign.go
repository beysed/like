package grammar

type Assign struct {
	Identifier MemberList
	Value      Expression
}

func (a Assign) String() string {
	return a.Identifier.String() + " = " + a.Value.String()
}

func (a Assign) Evaluate(context *Context) (any, error) {
	var v any
	var err error

	if a.Value != nil {
		v, err = a.Value.Evaluate(context)
	}

	if err == nil {
		context.Locals[a.Identifier[0].String()] = v
	}

	return v, err
}
