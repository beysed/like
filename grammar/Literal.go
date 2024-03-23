package grammar

func Literal(s string) Value {
	return Value{V: s}
}
