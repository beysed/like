package grammar

import "strings"

const ValueKey = "$value"

type Store map[string]any

type SystemContext struct {
	Buffer strings.Builder
}

func MakeSystemContext() SystemContext {
	return SystemContext{}
}

func (c *SystemContext) Output(text string) {
}

type InvokeResult struct {
	Stdout  string
	Stderr  string
	ErrCode int
}

func InvokeCommand(command string, args ...[]string) InvokeResult {
	return InvokeResult{}
}

type Context struct {
	Locals  Store
	Globals Store
	System  System
}
