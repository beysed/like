package grammar

type Assign struct {
	Identifier MemberList
	Value      Expression
}

func (a Assign) String() string {
	return a.Identifier.String() + " = " + a.Value.String()
}

func (a Assign) Evaluate(context *Context) (any, error) {
	store, _ := a.Identifier.Evaluate(context)
	v, _ := a.Value.Evaluate(context)

	store.(Store)["value"] = v

	return v, nil
}
