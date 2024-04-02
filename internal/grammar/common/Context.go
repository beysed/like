package common

type Context struct {
	Locals  Store
	Globals Store
	BuiltIn BuiltIn
	System  System
}
