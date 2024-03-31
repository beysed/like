package grammar

import (
	"fmt"
	"strings"
	"time"

	"github.com/beysed/shell/execute"
)

type Invoke struct {
	Expressions Expressions
}

func (a Invoke) String() string {
	return fmt.Sprintf("& %s", a.Expressions.String())
}

func (a Invoke) Evaluate(context *Context) (any, error) {
	output := strings.Builder{}

	cmdEval, err := a.Expressions[0].Evaluate(context)
	if err != nil {
		return cmdEval, err
	}

	cmd, ok := cmdEval.(string) // todo check
	if !ok {
		return cmdEval, MakeError("command is not string")
	}

	args := []string{}
	for _, v := range a.Expressions[1:] {
		arge, ok := v.(Expressions)
		if !ok || len(arge) != 2 {
			return v, MakeError("something wrong with arguments structure")
		}

		v, err := arge[1].Evaluate(context)
		if err != nil {
			return nil, err
		}

		if str, ok := v.(string); !ok {
			return v, MakeError("argument is not string")
		} else {
			args = append(args, str)
		}
	}

	command := execute.MakeCommand(cmd, args...)
	execution, _ := execute.Execute(command)

	close(execution.Stdin)

	for run := true; run; {
		select {
		case out := <-execution.Stdout:
			output.Write(out)
		case <-execution.Exit:
			run = false
		case <-time.After(time.Second * 10):
			execution.Kill()
			run = false
		}
	}

	return output.String(), nil
}
