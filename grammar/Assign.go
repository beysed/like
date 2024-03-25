package grammar

import . "like/expressions"

type Assign struct {
	Identifier string
	Value      Expression
}

func (a Assign) String() string {
	return a.Identifier + " = " + a.Value.String()
}

func (a Assign) Evaluate(system System, context *Context) (any, error) {
	var v any
	var err error

	if a.Value != nil {
		v, err = a.Value.Evaluate(system, context)
	}

	if err == nil {
		context.Locals[a.Identifier] = v
	}

	return v, err
}
