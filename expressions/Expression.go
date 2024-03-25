package expressions

type Expression interface {
	Evaluate(system System, context *Context) (any, error)
	String() string
}
