package common

type BuiltIn map[string]func(context *Context, args []any) (any, error)
