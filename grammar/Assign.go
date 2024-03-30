package grammar

type Assign struct {
	Store Expression
	Value Expression
}

func (a Assign) String() string {
	return a.Store.String() + " = " + a.Value.String()
}

func (a Assign) Evaluate(context *Context) (any, error) {
	store, _ := a.Store.Evaluate(context)
	v, _ := a.Value.Evaluate(context)

	store.(Store)[ValueKey] = v

	return v, nil
}
