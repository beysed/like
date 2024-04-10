package common

import (
	s "github.com/zeroflucs-given/generics/collections/stack"
)

type Context struct {
	Locals    Store
	Globals   Store
	BuiltIn   BuiltIn
	System    System
	PathStack *s.Stack[string]
}
