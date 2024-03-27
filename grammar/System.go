package grammar

type System interface {
	Output(text any)
	Invoke(command string, args ...[]string) (InvokeResult, error)
	// throw
	// input+
}
