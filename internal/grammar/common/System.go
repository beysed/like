package common

type System interface {
	OutputText(text string)
	OutputError(text string)
	ResolvePath(context *Context, filePath string) (string, error)
}
