package common

import (
	"strings"

	s "github.com/zeroflucs-given/generics/collections/stack"
)

type Locals struct {
	Store  Store
	Input  string
	Output strings.Builder
}

func MakeLocals(store Store) *Locals {
	return &Locals{
		Store: store,
	}
}

type Context struct {
	Locals    *s.Stack[*Locals]
	BuiltIn   BuiltIn
	System    System
	PathStack *s.Stack[string]
}
