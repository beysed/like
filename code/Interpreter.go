package code

import (
	// . "like/grammar"
	. "like/expressions"
)

type Interpreter struct {
	Context Context
}

type Environment map[string]string

type Executor interface {
	Execute(code []byte) (string, error)
}

// func (i *Interpreter) Execute(code []byte) (string, error) {
// }

// Output(text string)
// Invoke(command string, args ...[]string) (string, error)
