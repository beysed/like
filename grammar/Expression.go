package grammar

type Expression interface {
	Evaluate(system System, globals Context, locals Context) any
	String() string
}
