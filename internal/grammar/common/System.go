package common

type System interface {
	OutputText(text string)
	OutputError(text string)
	ResolvePath(context *Context, filePath string) (string, error)
	Invoke(executable string, args []string, stdin string) (string, string, error)
}
