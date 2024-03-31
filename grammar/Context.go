package grammar

import (
	"fmt"
	"strings"
)

const ValueKey = "$value"

type Store map[string]any

type SystemContext struct {
	Buffer strings.Builder
}

func MakeSystemContext() SystemContext {
	return SystemContext{}
}

func (c *SystemContext) Output(text string) {
	fmt.Print(text)
}

type Context struct {
	Locals  Store
	Globals Store
	System  System
}
