package grammar

type Expression interface {
	Evaluate(context *Context) (any, error)
	String() string
}
