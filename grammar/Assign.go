package grammar

type Assign struct {
	Identifier string
	Value      Expression
}

func (a Assign) String() string {
	return a.Identifier + " = " + a.Value.String()
}

func (a Assign) Evaluate(system System, globals Context, locals Context) any {
	var v any

	if a.Value != nil {
		v = a.Value.Evaluate(system, globals, locals)
	}

	locals[a.Identifier] = v
	return v
}
