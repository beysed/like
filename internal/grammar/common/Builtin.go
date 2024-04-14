package common

type BuiltIn map[string]func(context *Context, args []NamedValue) (any, error)
