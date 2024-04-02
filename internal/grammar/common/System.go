package common

type System interface {
	Output(text any)
	ResolvePath(filePath string) (string, error)
}
