package common

type System interface {
	Output(text any)
	ResolvePath(filePath string) (string, error)
	ReadFile(filePath string) ([]byte, error)
}
