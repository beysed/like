package common

import (
	s "github.com/zeroflucs-given/generics/collections/stack"
)

type Context struct {
	Locals    *s.Stack[Store]
	BuiltIn   BuiltIn
	System    System
	PathStack *s.Stack[string]
}
