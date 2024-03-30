package grammar

type Expression interface {
	String() string
	Evaluate(context *Context) (any, error)
}
