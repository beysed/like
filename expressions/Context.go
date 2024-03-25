package expressions

type Store map[string]any

type Context struct {
	Locals  Store
	Globals Store
	Builtin Store
}
