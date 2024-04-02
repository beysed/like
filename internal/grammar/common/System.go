package common

type System interface {
	Output(text any)
	ReadFile(filePath string) ([]byte, error)
	// throw
	// input+
}
