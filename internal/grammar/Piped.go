package grammar

type Piped interface {
	From() Ref[Expression]
	To() Ref[Expression]
}

type PipedInstance interface {
	Pipe | PipeIn | PipeOut | PipeAppend
	Expression
}

type PipedRef[T any] interface {
	*T
	Piped
}

type PipeOutInstance interface {
	*PipeOut | *PipeAppend
	Piped
}

func MakePiped[T PipedInstance, U PipedRef[T]](from Expression, to Expression) Expression {
	t := T{}
	var p U = &t
	p.From().Set(&from)
	p.To().Set(&to)

	return t
}
