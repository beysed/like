package expressions

type System interface {
	Output(text string)
	Invoke(command string, args ...[]string) (string, error)
	// throw
	// input
}
