package grammar

import (
	"fmt"

	c "github.com/beysed/like/internal/grammar/common"
)

type Pipe struct {
	from Expression
	to   Expression
	err  Expression
}

func (a *Pipe) From() Ref[Expression] {
	return MakeRef(&a.from)
}

func (a *Pipe) To() Ref[Expression] {
	return MakeRef(&a.to)
}

func (a *Pipe) Err() Ref[Expression] {
	return MakeRef(&a.err)
}

func (a Pipe) String() string {
	return fmt.Sprintf("%s | %s", a.from, a.to)
}

func makeAssign(store Expression, val any) Expression {
	assign := Assign{
		Store: store,
		Value: MakeConstant(val)}
	return assign
}

func isVal(expr Expression) bool {
	_, isRef := expr.(Reference)
	_, isLit := expr.(Literal)
	_, isConst := expr.(Constant)
	_, isString := expr.(ParsedString)
	_, isExpressions := expr.(Expressions)

	return isRef || isLit || isConst || isString || isExpressions
}

func (a Pipe) Evaluate(context *c.Context) (any, error) {
	resFrom, err := a.from.Evaluate(context)
	if err != nil {
		return a.from, err
	}

	var inputOut string
	var inputErr string

	isValFrom := isVal(a.from)

	_, current := context.Locals.Peek()

	if isValFrom {
		inputOut = c.Stringify(resFrom)
	} else {
		inputOut = current.Output.String()
		inputErr = current.Errors.String()
		current.Reset()
	}

	refTo, isRefTo := a.to.(Reference)
	refErr, isRefErr := a.err.(Reference)

	execPipe := func(input string, expr Expression) (any, error) {
		locals := c.MakeLocals(c.Store{})
		locals.Input = input
		context.Locals.Push(locals)

		res, err := expr.Evaluate(context)

		current.Output.WriteString(locals.Output.String())
		current.Errors.WriteString(locals.Errors.String())
		context.Locals.Pop()

		return res, err
	}

	var res any

	if isRefTo {
		assign := makeAssign(refTo.Expression, MakeConstant(inputOut))
		res, err = assign.Evaluate(context)
		if err != nil {
			return assign, err
		}
	} else {
		res, err = execPipe(inputOut, a.to)
		if err != nil {
			return a.to, err
		}
	}

	if a.err != nil {
		if isRefErr {
			assign := makeAssign(refErr.Expression, MakeConstant(inputErr))
			res, err = assign.Evaluate(context)
			if err != nil {
				return assign, err
			}
		} else {
			res, err = execPipe(inputErr, a.err)
			if err != nil {
				return a.to, err
			}
		}
	} else {
		context.System.OutputError(inputErr)
	}

	return res, err
}
